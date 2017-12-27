package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

// SongJSON is a struct representation of the SongJSON format
type SongJSON struct {
	Title        string `json:"title"`
	Lyrics       string `json:"lyrics"`
	Presentation string `json:"presentation,omitempty"`
	Author       string `json:"author,omitempty"`
	CCLI         int    `json:"ccli,omitempty"`
}

var lineMatcher, _ = regexp.Compile(`^(!|;|\.|\s)[a-zA-Z\s.\,\;]+`)

// SongMap is a structure to look up the song in-memory
type SongMap map[string]SongJSON

func main() {
	songMap := readSongs(os.Args[1])

	http.Handle("/songlist/", songListHandler(&songMap))
	http.Handle("/song/", songHandler(&songMap))
	http.Handle("/newSong", newSongHandler(&songMap))

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
		b, _ := ioutil.ReadFile(songDir + "/" + sf.Name())
		var newSong SongJSON
		json.Unmarshal(b, &newSong)

		outSongMap[newSong.Title] = newSong
	}

	return outSongMap
}

func songListHandler(smp *SongMap) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		songMap := *smp
		songList := []string{}

		for key := range songMap {
			songList = append(songList, key)
		}

		outJSON, _ := json.Marshal(songList)
		w.Write(outJSON)
	})
}

func songHandler(smp *SongMap) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		songMap := *smp
		songTitle := r.URL.Path[len("/song/"):]
		b, _ := json.Marshal(songMap[songTitle])
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})
}

func newSongHandler(smp *SongMap) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(scrubSong(r.Body)))
	})
}

func scrubSong(r io.Reader) string {
	b, _ := ioutil.ReadAll(r)
	s := string(b)
	// s := strings.Replace(string(b), "\r", "\n", -1)
	if lineMatcher.Match([]byte(s)) {
		return "The input is valid"
	}
	return s
}
