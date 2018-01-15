package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"

	"github.com/gorilla/schema"
)

// SongJSON is a struct representation of the SongJSON format
type SongJSON struct {
	Title        string `json:"title"`
	Body         string `json:"body"`
	Presentation string `json:"presentation,omitempty"`
	Author       string `json:"author,omitempty"`
	CCLI         int    `json:"ccli,omitempty"`
}

// SongMap is a structure to look up the song in-memory
type SongMap map[string]SongJSON

// Gorilla decoder
var decoder = schema.NewDecoder()

func main() {
	songMap := readSongs(os.Args[1])

	http.Handle("/songlist/", songListHandler(&songMap))
	http.Handle("/song/", songReader(&songMap))
	http.Handle("/newSong", createNewSong(&songMap, os.Args[1]))
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./frontend/src/static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		if r.URL.Path == "/" {
			f, _ := ioutil.ReadFile("./frontend/index.html")
			w.Write(f)
		} else {
			f, _ := ioutil.ReadFile("./frontend/" + r.URL.Path)
			w.Write(f)
		}
	})
	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func readSongs(songDir string) SongMap {
	songFiles, err := ioutil.ReadDir(songDir)
	if err != nil {
		panic(err)
	}

	outSongMap := SongMap{}

	for _, sf := range songFiles {
		fp := songDir + "/" + sf.Name()
		newSong := readSong(fp)

		outSongMap[newSong.Title] = newSong
	}

	return outSongMap
}

func readSong(filePath string) SongJSON {
	b, _ := ioutil.ReadFile(filePath)
	var newSong SongJSON
	json.Unmarshal(b, &newSong)

	return newSong
}

func songListHandler(smp *SongMap) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		songMap := *smp
		songList := []string{}

		for key := range songMap {
			songList = append(songList, key)
		}

		sort.Strings(songList)

		outJSON, _ := json.Marshal(songList)
		w.Write(outJSON)
	})
}

func songReader(smp *SongMap) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		songMap := *smp
		songTitle := r.URL.Path[len("/song/"):]
		b, _ := json.Marshal(songMap[songTitle])
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})
}

func createNewSong(smp *SongMap, songDir string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Scrub the data of any invalid or malicious input
		scrubbedTitle, titleErr := scrubUserTitle(r.PostFormValue("title"))
		scrubbedBody, bodyErr := scrubUserData(r.PostFormValue("body"))

		if titleErr != nil || bodyErr != nil {
			fmt.Fprintf(w, "Error: Invalid input")
			return
		}

		songMap := *smp

		contents := SongJSON{
			Title: scrubbedTitle,
			Body:  scrubbedBody,
		}

		songMap[scrubbedTitle] = contents

		serialized, _ := json.Marshal(contents)

		newFilePath := songDir + "/" + scrubbedTitle + ".json"
		ioutil.WriteFile(newFilePath, serialized, 0677)

		readSong(newFilePath)

		fmt.Fprintf(w, "Form submitted")
	})
}
