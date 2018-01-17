package main

import (
	"encoding/json"
	"io/ioutil"
)

// SongJSON is a struct representation of the SongJSON format
type SongJSON struct {
	Title        string `json:"title"`
	Body         string `json:"body"`
	Presentation string `json:"presentation,omitempty"`
	Author       string `json:"author,omitempty"`
	CCLI         int    `json:"ccli,omitempty"`
}

func readSongs(songDir string) Songs {
	songFiles, err := ioutil.ReadDir(songDir)
	if err != nil {
		panic(err)
	}

	outSongs := Songs{}

	for _, sf := range songFiles {
		fp := songDir + "/" + sf.Name()
		newSong := readSong(fp)

		outSongs[newSong.Title] = newSong
	}

	return outSongs
}

func readSong(filePath string) SongJSON {
	b, _ := ioutil.ReadFile(filePath)
	var newSong SongJSON
	json.Unmarshal(b, &newSong)

	return newSong
}
