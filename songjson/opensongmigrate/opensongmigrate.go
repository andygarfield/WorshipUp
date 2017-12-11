package opensongmigrate

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/andygarfield/hypher"
)

func main() {
	var songDir string
	if len(os.Args) < 2 {
		fmt.Println("Please enter directory where OpenSong songs are stored.")
	} else {
		songDir = os.Args[1]
	}
	os.Mkdir("./json", 0777)

	dirFiles, dirErr := ioutil.ReadDir(songDir)
	if dirErr != nil {
		fmt.Println("Could not read directory")
	}

	for _, fi := range dirFiles {
		b, err := ioutil.ReadFile(songDir + "/" + fi.Name())
		if err != nil {
			fmt.Println(err)
		}
		c, convertErr := convertOpenSong(bytes.NewReader(b))

		if convertErr == nil {
			ioutil.WriteFile("./json/"+fi.Name()+".json", c, 0777)
		} else {
			fmt.Println(fi.Name() + " could not be converted")
		}
	}
}

func convertOpenSong(r io.Reader) ([]byte, error) {
	addIfExists := func(m *map[string]string, contents []byte, tag string) error {
		p, err := hypher.FindTags(tag, bytes.NewReader(contents))
		if err == nil && p[0].Contents != "" {
			dm := *m
			dm[tag] = p[0].Contents
			return nil
		}

		return errors.New(tag + " tag was not found")
	}

	contents, _ := ioutil.ReadAll(r)
	song := map[string]string{}

	// Get and transform the lyrics section
	l, lyricsErr := hypher.FindTags("lyrics", bytes.NewReader(contents))
	if lyricsErr != nil || len(l) == 0 {
		return nil, errors.New("File had no lyrics tag")
	}

	l[0].Contents = strings.Replace(l[0].Contents, "\r", "\n", -1)
	ld := strings.Split(l[0].Contents, "\n")

	var convertedLyrics string
	for _, line := range ld {
		rs := []rune(line)

		if len(rs) > 0 && rs[0] == '[' {
			var section string
			for _, r := range rs[1:] {
				if r != ']' {
					section += string(r)
				} else {
					break
				}
			}

			convertedLyrics += "!" + section + "\n"
		} else {
			convertedLyrics += string(rs) + "\n"
		}
	}
	song["lyrics"] = convertedLyrics

	// Get the other values that don't need much conversion

	titleErr := addIfExists(&song, contents, "title")
	if titleErr != nil {
		return nil, titleErr
	}

	addIfExists(&song, contents, "presentation")
	addIfExists(&song, contents, "author")
	addIfExists(&song, contents, "ccli")

	marshalled, _ := json.Marshal(song)
	return marshalled, nil
}
