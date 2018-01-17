package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"time"
)

// Songs is a structure to look up a song data in memory
type Songs map[string]SongJSON

// Services is a structure to look up a service according to date
type Services map[time.Time][]SongJSON

func main() {
	songs := readSongs(os.Args[1])
	// serviceMap

	http.Handle("/songlist/", songListHandler(&songs))
	http.Handle("/song/", songReader(&songs))
	http.Handle("/songsubmit", submitSong(&songs, os.Args[1]))
	http.Handle("/", http.FileServer(http.Dir("./frontend")))

	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func songListHandler(sp *Songs) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		songs := *sp
		songList := []string{}

		for key := range songs {
			songList = append(songList, key)
		}

		sort.Strings(songList)

		outJSON, _ := json.Marshal(songList)
		w.Write(outJSON)
	})
}

func songReader(sp *Songs) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		songs := *sp
		songTitle := r.URL.Path[len("/song/"):]
		b, _ := json.Marshal(songs[songTitle])
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})
}

func submitSong(sp *Songs, songDir string) http.Handler {
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

		songs := *sp

		contents := SongJSON{
			Title: scrubbedTitle,
			Body:  scrubbedBody,
		}

		songs[scrubbedTitle] = contents

		serialized, _ := json.Marshal(contents)

		writeFilePath := songDir + "/" + scrubbedTitle + ".json"
		ioutil.WriteFile(writeFilePath, serialized, 0677)
		fmt.Fprintf(w, "Form submitted")
	})
}
