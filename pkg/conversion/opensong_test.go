package conversion

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestConvertOpenSongSet(t *testing.T) {
	readers := loopThroughDir("./testdata/OpenSongSet")

	var (
		sb []byte
	)

	for _, r := range readers {
		sb, _ = ioutil.ReadAll(r)
		oss := openSongSet(sb)
		oss.Convert()
	}
}

func TestConvertOpenSongSong(t *testing.T) {
	readers := loopThroughDir("./testdata/OpenSongSong")

	var (
		sb []byte
	)

	for _, r := range readers {
		sb, _ = ioutil.ReadAll(r)
		oss := openSongSong(sb)
		_, err := oss.Convert()
		if err != nil {
			fmt.Println(err)
		}

		// fmt.Println(song)
	}
}
