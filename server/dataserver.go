package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/NYTimes/gziphandler"

	"github.com/andygarfield/worshipup/pkg/worshipup"
	"github.com/boltdb/bolt"
)

func main() {
	// Create database and add buckets if they don't exist
	db, err := bolt.Open("library.db", 0600, nil)

	if err != nil {
		log.Fatalf("Open bolt database: %s", err)
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Songs"))
		if err != nil {
			return fmt.Errorf("Create songs bucket: %s", err)
		}
		_, err = tx.CreateBucketIfNotExists([]byte("Sets"))
		if err != nil {
			return fmt.Errorf("Create sets bucket: %s", err)
		}
		return nil
	})

	http.Handle("/songlist/", gziphandler.GzipHandler(getSongList(db)))
	http.Handle("/song/", songHandler(db))
	http.Handle("/songsubmit", submitSong(db))
	http.Handle("/", gziphandler.GzipHandler(http.FileServer(http.Dir("./frontend"))))

	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func getSongList(db *bolt.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		songList := []string{}

		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("Songs"))
			c := b.Cursor()

			for k, _ := c.First(); k != nil; k, _ = c.Next() {
				songList = append(songList, string(k))
			}
			return nil
		})

		sort.Strings(songList)

		outJSON, _ := json.Marshal(songList)
		w.Write(outJSON)
	})
}

func songHandler(db *bolt.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			songTitle := r.URL.Path[len("/song/"):]
			var songData []byte

			db.View(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte("Songs"))
				songData = b.Get([]byte(songTitle))
				return nil
			})

			w.Header().Set("Content-Type", "application/json")
			w.Write(songData)
		} else if r.Method == "DELETE" {
			db.Update(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte("Songs"))
				b.Delete([]byte(r.URL.Path[len("/song/"):]))
				return nil
			})
		}
	})
}

func submitSong(db *bolt.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Scrub the data of any invalid or malicious input
		scrubbedTitle, titleErr := worshipup.ScrubUserTitle(r.PostFormValue("title"))
		scrubbedBody, bodyErr := worshipup.ScrubUserData(r.PostFormValue("body"))

		if titleErr != nil {
			fmt.Fprintf(w, fmt.Sprint(titleErr))
			return
		}
		if bodyErr != nil {
			fmt.Fprintf(w, fmt.Sprint(bodyErr))
			return
		}

		contents := worshipup.SongJSON{
			Title: scrubbedTitle,
			Body:  scrubbedBody,
		}

		serialized, _ := json.Marshal(contents)

		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("Songs"))
			b.Put([]byte(scrubbedTitle), serialized)
			return nil
		})
		fmt.Fprintf(w, "Form submitted")
	})
}
