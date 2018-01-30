package conversion

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/andygarfield/worshipup/pkg/worshipup"
	"github.com/boltdb/bolt"
)

// SetConverter takes the input set and converts it to a worshipup.ServiceOrder
type SetConverter interface {
	Convert() (worshipup.SetOrder, error)
}

// SongConverter takes the input song and converts it to a worshipup.SongJSON
type SongConverter interface {
	Convert() (worshipup.SongJSON, error)
}

// ImportSet takes a ServiceConverter and imports it into the app's database in the "Sets" bucket
func ImportSet(db *bolt.DB, s SetConverter) {
	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte("Sets"))
		set, err := s.Convert()
		if err != nil {
			fmt.Println(err)
		}
		m, err := json.Marshal(set)
		if err != nil {
			fmt.Println(err)
		}

		date := set.Date.Format("20060102")
		err = b.Put([]byte(date), m)

		return nil
	})
}

// ImportSong takes a SongConverter and imports it into the app's database in the "Songs" bucket
func ImportSong(db *bolt.DB, s SongConverter) {
	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte("Songs"))
		song, err := s.Convert()
		if err != nil {
			fmt.Println(err)
		}
		m, err := json.Marshal(song)
		if err != nil {
			fmt.Println(err)
		}

		err = b.Put([]byte(song.Title), m)

		return nil
	})
}

// ImportSetDir imports all of the sets in a directory into the app's database
func ImportSetDir(db *bolt.DB, dir, setType string) error {
	readers := loopThroughDir(dir)
	if strings.ToLower(setType) == "opensong" {
		for _, r := range readers {
			b, _ := ioutil.ReadAll(r)

			ImportSet(db, openSongSet(b))
		}
	} else {
		return fmt.Errorf("%s is not implemented as an import type at this time", setType)
	}

	return nil
}

// ImportSongDir imports all of the songs in a directory into the app's database
func ImportSongDir(db *bolt.DB, dir, songType string) error {
	readers := loopThroughDir(dir)
	if strings.ToLower(songType) == "opensong" {
		for _, r := range readers {
			b, _ := ioutil.ReadAll(r)

			ImportSong(db, openSongSong(b))
		}
	} else {
		return fmt.Errorf("%s is not implimented as in import type at this time", songType)
	}

	return nil
}

func loopThroughDir(dirname string) []io.Reader {
	fis, err := ioutil.ReadDir(dirname)
	if err != nil {
		fmt.Println("Test directory not found")
	}

	readers := []io.Reader{}
	for _, fi := range fis {
		f, _ := os.Open(filepath.Join(dirname, fi.Name()))
		readers = append(readers, f)
	}

	return readers
}
