package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
)

// SongJSON is a struct representation of the SongJSON format
type SongJSON struct {
	Title        string `json:"title"`
	Lyrics       string `json:"lyrics"`
	Presentation string `json:"presentation,omitempty"`
	Author       string `json:"author,omitempty"`
	CCLI         int    `json:"ccli,omitempty"`
}

// SongMap is a structure to look up the song in-memory
type SongMap map[string]SongJSON

func main() {
	songMap := readSongs(os.Args[1])

	http.Handle("/songlist/", songListHandler(&songMap))
	http.Handle("/song/", songReader(&songMap))
	http.Handle("/newSong", createNewSong(&songMap, os.Args[1]))

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
		scrubbedTitle, titleErr := scrubUserTitle(r.PostFormValue("title"))
		scrubbedBody, bodyErr := scrubUserData(r.PostFormValue("body"))

		if titleErr != nil || bodyErr != nil {
			fmt.Fprintf(w, "Error: Invalid input")
		}

		songMap := *smp
		songMap[scrubbedTitle] = SongJSON{
			Title:  scrubbedTitle,
			Lyrics: scrubbedBody,
		}

		newFilePath := songDir + "/" + scrubbedTitle + ".json"
		ioutil.WriteFile(newFilePath, []byte(scrubbedBody), 0677)

		readSong(newFilePath)

		fmt.Fprintf(w, scrubbedBody)
	})
}
