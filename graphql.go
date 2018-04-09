package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/andygarfield/worshipup/pkg/types"
	"github.com/andygarfield/worshipup/pkg/utils"
	"github.com/boltdb/bolt"
	graphql "github.com/graph-gophers/graphql-go"
)

// Resolver is the root resolver
type Resolver struct{}

// Song type
// Song returns a song from the database
func (r *Resolver) Song(args struct{ ID graphql.ID }) (*songResolver, error) {
	song := types.SongJSON{}
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Songs"))
		s := utils.FetchID(args.ID, b)
		json.Unmarshal(s, &song)
		return nil
	})

	return &songResolver{&song}, nil
}

// AllSongs fetches all the songs in the database
func (r *Resolver) AllSongs() (*[]*songResolver, error) {
	sr := []*songResolver{}
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Songs"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			s := types.SongJSON{}
			json.Unmarshal(v, &s)
			sr = append(sr, &songResolver{&s})
		}
		return nil
	})
	return &sr, nil
}

// AddSong adds a new song to the database after scrubbing the input
func (r *Resolver) AddSong(args *struct{ Song types.SongJSON }) *songResolver {
	// Replacing new lines with ~ because github.com/graph-gophers/graphql-go
	// is currently using go's text/scanner, which doesn't handle new lines
	// correctly
	args.Song.Body = strings.Replace(args.Song.Body, "~", "\n", -1)

	err := args.Song.Scrub()
	if err != nil {
		fmt.Println(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Songs"))
		id, _ := b.NextSequence()
		args.Song.ID = id

		sb, err := json.Marshal(args.Song)
		if err != nil {
			fmt.Println(err)
		}

		err = b.Put(utils.EncodeUint64(id), sb)
		return err
	})
	if err != nil {
		fmt.Println(err)
	}
	return &songResolver{&args.Song}
}

// UpdateSong changes a song in the database after scrubbing the input
func (r *Resolver) UpdateSong(args *struct {
	ID   graphql.ID
	Song types.SongJSON
}) *songResolver {
	// Replacing new lines with ~ because github.com/graph-gophers/graphql-go
	// is currently using go's text/scanner, which doesn't handle new lines
	// correctly
	args.Song.Body = strings.Replace(args.Song.Body, "~", "\n", -1)

	err := args.Song.Scrub()
	if err != nil {
		fmt.Println(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Songs"))

		id, err := strconv.ParseUint(string(args.ID), 10, 64)
		args.Song.ID = id

		sb, err := json.Marshal(args.Song)
		if err != nil {
			fmt.Println(err)
		}

		err = b.Put(utils.EncodeUint64(id), sb)
		return err
	})
	if err != nil {
		fmt.Println(err)
	}
	return &songResolver{&args.Song}
}

// DeleteSong deletes a song from the database
func (r *Resolver) DeleteSong(args *struct{ ID graphql.ID }) *string {
	var s string
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Songs"))
		b.Delete(utils.IDToUintBytes(args.ID))
		return nil
	})
	if err != nil {
		s = fmt.Sprint(err)
		return &s
	}
	s = fmt.Sprintf("Deleted song (ID:%v) successfully", args.ID)
	return &s
}

type songResolver struct {
	s *types.SongJSON
}

func (r *songResolver) ID() graphql.ID {
	return graphql.ID(fmt.Sprint(r.s.ID))
}

func (r *songResolver) Title() string {
	return r.s.Title
}

func (r *songResolver) Body() string {
	return r.s.Body
}

func (r *songResolver) Presentation() *string {
	return r.s.Presentation
}

func (r *songResolver) Author() *string {
	return r.s.Author
}

func (r *songResolver) CCLI() *int32 {
	return r.s.CCLI
}

// Set type
// Set returns a set from the database
func (r *Resolver) Set(args struct{ ID graphql.ID }) (*setResolver, error) {
	set := types.SetList{}
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Sets"))
		s := utils.FetchID(args.ID, b)
		json.Unmarshal(s, &set)
		return nil
	})

	return &setResolver{&set}, nil
}

// AllSets returns all the sets in the database
func (r *Resolver) AllSets() (*[]*setResolver, error) {
	sr := []*setResolver{}
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Sets"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			s := types.SetList{}
			json.Unmarshal(v, &s)
			sr = append(sr, &setResolver{&s})
		}
		return nil
	})
	return &sr, nil
}

// AddSet adds a blank set to the database
func (r *Resolver) AddSet(args *struct{ Set types.SetList }) *setResolver {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Sets"))
		id, _ := b.NextSequence()
		args.Set.ID = id

		sb, err := json.Marshal(args.Set)
		if err != nil {
			fmt.Println(err)
		}

		err = b.Put(utils.EncodeUint64(id), sb)
		return err
	})
	if err != nil {
		fmt.Println(err)
	}
	return &setResolver{&args.Set}
}

// UpdateSet changes the set in the database
func (r *Resolver) UpdateSet(args *struct {
	ID  graphql.ID
	Set types.SetList
}) *setResolver {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Sets"))

		id, err := strconv.ParseUint(string(args.ID), 10, 64)
		args.Set.ID = id

		sb, err := json.Marshal(args.Set)
		if err != nil {
			fmt.Println(err)
		}

		err = b.Put(utils.EncodeUint64(id), sb)
		return err
	})
	if err != nil {
		fmt.Println(err)
	}
	return &setResolver{&args.Set}
}

// DeleteSet deletes a set from the database
func (r *Resolver) DeleteSet(args *struct{ ID graphql.ID }) *string {
	var s string
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Sets"))
		b.Delete(utils.IDToUintBytes(args.ID))
		return nil
	})
	if err != nil {
		s = fmt.Sprint(err)
		return &s
	}
	s = fmt.Sprintf("Deleted set (ID:%v) successfully", args.ID)
	return &s
}

type setResolver struct {
	s *types.SetList
}

func (r *setResolver) ID() graphql.ID {
	return graphql.ID(fmt.Sprint(r.s.ID))
}

func (r *setResolver) Date() int32 {
	return int32(r.s.Date.Unix())
}

func (r *setResolver) Songs() *[]*songResolver {
	l := make([]*songResolver, len(*r.s.Songs))
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Songs"))
		for i, id := range *r.s.Songs {
			song := types.SongJSON{}
			json.Unmarshal(b.Get(utils.EncodeUint64(*id)), &song)
			l[i] = &songResolver{&song}
		}
		return nil
	})
	return &l
}
