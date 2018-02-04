package conversion

import (
	"bytes"
	"strconv"
	"strings"
	"time"

	"github.com/andygarfield/worshipup/pkg/worshipup"
	"github.com/antchfx/xmlquery"
)

// OpenSongSet is an XML encoded set that is used by OpenSong
type OpenSongSet []byte

// OpenSongSong is an XML encoded song that is used by OpenSong
type OpenSongSong []byte

// Convert converts an OpenSongSet into a worshipup.ServiceOrder
func (set OpenSongSet) Convert() (worshipup.SetOrder, error) {
	b := bytes.NewBuffer(set)
	doc, err := xmlquery.Parse(b)
	if err != nil {
		return worshipup.SetOrder{}, err
	}

	var (
		date  time.Time
		songs []string
	)

	// Find date
	for _, el := range xmlquery.Find(doc, `/set`) {
		for _, a := range el.Attr {
			if a.Name.Local == "name" {
				date, err = time.Parse("Worship2006_01_02", a.Value)
				if err != nil {
					return worshipup.SetOrder{}, err
				}
			}
		}
	}

	// Find songs
	for _, el := range xmlquery.Find(doc, `//slide_group[@type='song']`) {
		for _, a := range el.Attr {
			if a.Name.Local == "name" {
				songs = append(songs, a.Value)
			}
		}
	}

	return worshipup.SetOrder{
		Date:  date,
		Songs: songs,
	}, nil
}

// Convert converts an OpenSongSong into a worshipup.SongJSON
func (song OpenSongSong) Convert() (worshipup.SongJSON, error) {
	getTagInnerText := func(doc *xmlquery.Node, tag string) string {
		return xmlquery.Find(doc, "//"+tag)[0].InnerText()
	}

	b := bytes.NewBuffer(song)
	doc, err := xmlquery.Parse(b)
	if err != nil {
		return worshipup.SongJSON{}, err
	}
	// Get and transform song body
	body := getTagInnerText(doc, "lyrics")

	body = strings.Replace(body, "\r", "\n", -1)
	ld := strings.Split(body, "\n")

	var convertedBody string

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
			convertedBody += "!" + section + "\n"
		} else if len(rs) >= 3 && line[:3] == " ||" {
			convertedBody += "\n"
		} else if len(rs) >= 3 && line[:3] == "---" {
			convertedBody += "\n"
		} else {
			convertedBody += string(rs) + "\n"
		}
	}

	// After conversion, scrub data for any illegal characters
	convertedBody, err = worshipup.ScrubUserData(convertedBody)
	if err != nil {
		return worshipup.SongJSON{}, err
	}

	// Get the other values that don't need much conversion
	title := getTagInnerText(doc, "title")
	presentation := getTagInnerText(doc, "presentation")
	author := getTagInnerText(doc, "author")
	ccli, err := strconv.Atoi(getTagInnerText(doc, "ccli"))
	if err != nil {
		ccli = -1
	}

	return worshipup.SongJSON{
		Title:        title,
		Body:         convertedBody,
		Presentation: presentation,
		Author:       author,
		CCLI:         ccli,
	}, nil
}
