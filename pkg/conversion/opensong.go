package conversion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/andygarfield/worshipup/pkg/types"
	"github.com/andygarfield/worshipup/pkg/utils"
	"github.com/antchfx/xmlquery"
	"github.com/boltdb/bolt"
)

// OpenSongSet is an XML encoded set that is used by OpenSong
type OpenSongSet []byte

// OpenSongSong is an XML encoded song that is used by OpenSong
type OpenSongSong []byte

// Convert converts an OpenSongSet into a worshipup.ServiceList
func (set OpenSongSet) Convert(db *bolt.DB) (types.SetList, error) {
	b := bytes.NewBuffer(set)
	doc, err := xmlquery.Parse(b)
	if err != nil {
		return types.SetList{}, err
	}

	var date time.Time
	songs := []*uint64{}

	// Find date
	for _, el := range xmlquery.Find(doc, `/set`) {
		for _, a := range el.Attr {
			if a.Name.Local == "name" {
				date, err = time.Parse("Worship2006_01_02", a.Value)
				if err != nil {
					return types.SetList{}, err
				}
			}
		}
	}

	// Use song titles to find the matching song and populate the song slice
	for _, el := range xmlquery.Find(doc, `//slide_group[@type='song']`) {
		for _, a := range el.Attr {
			if a.Name.Local == "name" {
				// Fetch matching IDs from bolt
				db.View(func(tx *bolt.Tx) error {
					b := tx.Bucket([]byte("Songs"))

					c := b.Cursor()
					for k, v := c.First(); k != nil; k, v = c.Next() {
						song := types.SongJSON{}
						json.Unmarshal(v, &song)
						if a.Value == song.Title {
							newInt := utils.DecodeUint64(k)
							songs = append(songs, &newInt)
						}
					}

					return nil
				})
			}
		}
	}

	return types.SetList{
		Date:  date,
		Songs: &songs,
	}, nil
}

// Convert converts an OpenSongSong into a worshipup.SongJSON
func (sb OpenSongSong) Convert() (types.SongJSON, error) {
	getTagInnerText := func(doc *xmlquery.Node, tag string) (string, error) {
		results := xmlquery.Find(doc, "//"+tag)
		if len(results) == 0 {
			return "", fmt.Errorf("Could not find %s tag", tag)
		}
		return results[0].InnerText(), nil
	}

	b := bytes.NewBuffer(sb)
	doc, err := xmlquery.Parse(b)
	if err != nil {
		return types.SongJSON{}, err
	}

	song := types.SongJSON{}

	// Get data from OpenSong file
	song.Body, err = getTagInnerText(doc, "lyrics")
	if err != nil {
		fmt.Println(err)
	}
	song.Title, err = getTagInnerText(doc, "title")
	if err != nil {
		fmt.Println(err)
	}

	presentation, err := getTagInnerText(doc, "presentation")
	if err != nil {
		fmt.Println(err)
	}
	song.Presentation = &presentation

	author, err := getTagInnerText(doc, "author")
	if err != nil {
		fmt.Println(err)
	}
	song.Author = &author

	// Convert CCLI around
	ccli, err := getTagInnerText(doc, "ccli")
	if err != nil {
		fmt.Println(err)
	}
	ccliInt, err := strconv.ParseInt(ccli, 10, 32)
	if err != nil {
		ccliInt = -1
	}
	ccliVal := int32(ccliInt)
	song.CCLI = &ccliVal

	// Normalize song.Body data
	song.Body = strings.Replace(song.Body, "\r\n", "\n", -1)
	ld := strings.Split(song.Body, "\n")

	song.Body = ""

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
			song.Body += "!" + section + "\n"
		} else if len(rs) >= 3 && line[:3] == " ||" {
			song.Body += "\n"
		} else if len(rs) >= 3 && line[:3] == "---" {
			song.Body += "\n"
		} else {
			song.Body += string(rs) + "\n"
		}
	}

	// After conversion, scrub data for any illegal characters
	err = song.Scrub()
	if err != nil {
		return types.SongJSON{}, err
	}

	return song, nil
}
