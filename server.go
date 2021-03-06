package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/andygarfield/worshipup/pkg/conversion"

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

	http.Handle("/songlist/", getSongList(db))
	http.Handle("/setlists/", getSetLists(db))
	http.Handle("/song/", songHandler(db))
	http.Handle("/songsubmit/", submitSong(db))
	http.Handle("/setsubmit/", submitSet(db))
	http.Handle("/songupload/", uploadSongs(db))
	http.Handle("/setupload/", uploadSets(db))
	http.Handle("/", gziphandler.GzipHandler(http.FileServer(http.Dir("./frontend"))))

	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func getSongList(db *bolt.DB) http.Handler {
	return gziphandler.GzipHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	}))
}

func getSetLists(db *bolt.DB) http.Handler {
	return gziphandler.GzipHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setLists := map[string]worshipup.SetList{}

		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("Sets"))
			c := b.Cursor()

			for _, setData := c.First(); setData != nil; _, setData = c.Next() {
				sl := worshipup.SetList{}
				json.Unmarshal(setData, &sl)
				dateString := sl.Date.Format("20060102")
				setLists[dateString] = sl
			}

			return nil
		})

		outJSON, _ := json.Marshal(setLists)
		w.Write(outJSON)
	}))
}

func songHandler(db *bolt.DB) http.Handler {
	return gziphandler.GzipHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	}))
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

func submitSet(db *bolt.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setOrder := strings.Split(r.PostFormValue("set"), ",")
		date, _ := time.Parse("20060102", r.PostFormValue("date"))

		sj := worshipup.SetList{
			Date:  date,
			Songs: setOrder,
		}

		slJSON, _ := json.Marshal(sj)

		fmt.Println(setOrder)
		err := db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("Sets"))
			err := b.Put([]byte(r.PostFormValue("date")), slJSON)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			w.Header().Set("Status Code", "500")
			fmt.Fprintf(w, fmt.Sprint(err))
			return
		}

		fmt.Fprintf(w, "Submitted")
	})
}

func uploadSongs(db *bolt.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(32 << 20)

		m := r.MultipartForm
		files := m.File["uploadfiles"]

		for i := range files {
			file, err := files[i].Open()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			b, _ := ioutil.ReadAll(file)
			song := conversion.OpenSongSong(b)

			err = conversion.ImportSong(db, song)
			if err != nil {
				fmt.Fprintf(w, fmt.Sprint(err))
			}

			defer file.Close()
		}
	})
}

func uploadSets(db *bolt.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(32 << 20)

		m := r.MultipartForm
		files := m.File["uploadfiles"]
		for i := range files {
			file, err := files[i].Open()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			b, _ := ioutil.ReadAll(file)
			set := conversion.OpenSongSet(b)

			err = conversion.ImportSet(db, set)
			if err != nil {
				fmt.Fprintf(w, fmt.Sprint(err))
			}

			defer file.Close()
		}
	})
}
