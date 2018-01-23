package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/boltdb/bolt"
)

// Songs is a structure to look up a song data in memory
type Songs map[string]SongJSON

// Services is a structure to look up a service according to date
type Services map[time.Time][]SongJSON

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

	http.Handle("/songlist/", getSongList(db))
	http.Handle("/song/", songReader(db))
	http.Handle("/songsubmit", submitSong(db))
	http.Handle("/", http.FileServer(http.Dir("./frontend")))

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

func songReader(db *bolt.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		songTitle := r.URL.Path[len("/song/"):]
		var songData []byte

		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("Songs"))
			songData = b.Get([]byte(songTitle))
			return nil
		})

		w.Header().Set("Content-Type", "application/json")
		w.Write(songData)
	})
}

func submitSong(db *bolt.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Scrub the data of any invalid or malicious input
		scrubbedTitle, titleErr := scrubUserTitle(r.PostFormValue("title"))
		scrubbedBody, bodyErr := scrubUserData(r.PostFormValue("body"))

		if titleErr != nil {
			fmt.Fprintf(w, fmt.Sprint(titleErr))
			return
		}
		if bodyErr != nil {
			fmt.Fprintf(w, fmt.Sprint(bodyErr))
			return
		}

		contents := SongJSON{
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

func getServiceList(s *Services) {

}
