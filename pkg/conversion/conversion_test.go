package conversion

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/boltdb/bolt"
)

func TestImportSet(t *testing.T) {
	testDB := "./test.db"

	db, _ := bolt.Open(testDB, 0600, nil)
	s, err := ioutil.ReadFile("./testdata/OpenSongSet/Worship2016_09_18")
	if err != nil {
		fmt.Println(err)
	}
	ImportSet(db, openSongSet(s))

	var data []byte
	db.View(func(tx *bolt.Tx) error {
		oss, err := openSongSet(s).Convert()
		if err != nil {
			return err
		}
		date := oss.Date.Format("20060102")

		b := tx.Bucket([]byte("Sets"))
		data = b.Get([]byte(date))

		return nil
	})

	fmt.Println(string(data))

	os.Remove(testDB)
}
