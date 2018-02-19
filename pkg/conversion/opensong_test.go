package conversion

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/andygarfield/worshipup/pkg/worshipup"
)

func TestConvertOpenSongSet(t *testing.T) {
	readers := loopThroughDir("./testdata/OpenSongSet")

	var (
		sb []byte
	)

	wus := worshipup.SetList{}

	for _, r := range readers[:1] {
		sb, _ = ioutil.ReadAll(r)
		oss := OpenSongSet(sb)
		wus, _ = oss.Convert()
	}

	fmt.Println(wus)
}

// func TestConvertOpenSongSong(t *testing.T) {
// 	readers := loopThroughDir("./testdata/OpenSongSong")

// 	var (
// 		sb []byte
// 	)

// 	for _, r := range readers {
// 		sb, _ = ioutil.ReadAll(r)
// 		oss := OpenSongSong(sb)
// 		_, err := oss.Convert()
// 		if err != nil {
// 			fmt.Println(err)
// 		}

// 		// fmt.Println(song)
// 	}
// }
