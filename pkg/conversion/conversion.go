package conversion

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/andygarfield/worshipup/pkg/types"
	"github.com/andygarfield/worshipup/pkg/utils"
	"github.com/boltdb/bolt"
)

// SetConverter takes the input set and converts it to a worshipup.ServiceList
// This songs in the set are required to have the song exist in the database,
// otherwise Convert() fails
type SetConverter interface {
	Convert(db *bolt.DB) (types.SetList, error)
}

// SongConverter takes the input song and converts it to a worshipup.SongJSON
type SongConverter interface {
	Convert() (types.SongJSON, error)
}

// ImportSet takes a ServiceConverter and imports it into the app's database
// in the "Sets" bucket
func ImportSet(db *bolt.DB, s SetConverter) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Sets"))
		set, err := s.Convert(db)
		if err != nil {
			return err
		}

		id, _ := b.NextSequence()
		set.ID = id

		m, err := json.Marshal(set)
		if err != nil {
			return err
		}

		err = b.Put(utils.EncodeUint64(id), m)

		return nil
	})

	return err
}

// ImportSong takes a SongConverter and imports it into the app's database in
// the "Songs" bucket
func ImportSong(db *bolt.DB, s SongConverter) error {
	// fmt.Println("Got here.")
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Songs"))
		song, err := s.Convert()
		if err != nil {
			return fmt.Errorf("Problem with conversion: %v", err)
		}

		id, _ := b.NextSequence()
		song.ID = id

		m, err := json.Marshal(song)
		if err != nil {
			return err
		}

		err = b.Put(utils.EncodeUint64(id), m)

		return nil
	})

	return err
}

// ImportSetDir imports all of the sets in a directory into the app's database
func ImportSetDir(db *bolt.DB, dir, setType string) error {
	readers := loopThroughDir(dir)
	if strings.ToLower(setType) == "opensong" {
		for _, r := range readers {
			b, _ := ioutil.ReadAll(r)

			ImportSet(db, OpenSongSet(b))
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
			b, err := ioutil.ReadAll(r)
			if err != nil {
				fmt.Println(err)
			}
			err = ImportSong(db, OpenSongSong(b))
			if err != nil {
				fmt.Println(err)
			}
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
