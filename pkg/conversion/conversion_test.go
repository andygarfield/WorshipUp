package conversion

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/boltdb/bolt"
)

// func TestImportSetDir(t *testing.T) {
// 	testDB := "../../library.db"

// 	db, _ := bolt.Open(testDB, 0600, nil)
// 	ImportSetDir(db, "testdata/OpenSongSet", "opensong")
// }

func TestImportSet(t *testing.T) {
	testDB := "../../library.db"
	db, _ := bolt.Open(testDB, 0600, nil)

	b, _ := ioutil.ReadFile("testdata/OpenSongSet/Worship2016_09_18")
	osSet := OpenSongSet(b)

	ImportSet(db, osSet)
	readBoltSets(db)
}

func readBoltSets(db *bolt.DB) {
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Sets"))

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}

		return nil
	})
}
