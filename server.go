package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/NYTimes/gziphandler"
	"github.com/andygarfield/worshipup/pkg/conversion"

	"github.com/boltdb/bolt"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

// Create database and add buckets if they don't exist
var db = dbInit()

func main() {
	// Create graphql schema
	schema := fetchSchema("/Users/andy/go/src/github.com/andygarfield/worshipup/schema.graphql")

	http.Handle("/", gziphandler.GzipHandler(http.FileServer(http.Dir("./frontend"))))
	http.Handle("/gql", &relay.Handler{Schema: schema})

	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dbInit() *bolt.DB {
	dbPath := "library.db"
	// If the DB doesn't exist, create it, and import songs at the end
	var db *bolt.DB

	_, err := os.Stat(dbPath)
	if err == nil {
		db, err := bolt.Open(dbPath, 0600, nil)
		if err != nil {
			log.Fatalf("Open bolt database: %s", err)
		}
		return db
	}

	db, err = bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatalf("Create bolt database: %s", err)
	}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("Songs"))
		if err != nil {
			return fmt.Errorf("Create songs bucket: %s", err)
		}
		_, err = tx.CreateBucket([]byte("Sets"))
		if err != nil {
			return fmt.Errorf("Create sets bucket: %s", err)
		}
		return nil
	})

	err = conversion.ImportSongDir(db, "./pkg/conversion/testdata/OpenSongSong", "opensong")
	if err != nil {
		fmt.Println(err)
	}
	err = conversion.ImportSetDir(db, "./pkg/conversion/testdata/OpenSongSet", "opensong")
	if err != nil {
		fmt.Println(err)
	}

	return db
}

func fetchSchema(schemaPath string) *graphql.Schema {
	s, err := ioutil.ReadFile(schemaPath)
	if err != nil {
		panic("Couldn't find schema file")
	}
	return graphql.MustParseSchema(string(s), &Resolver{})
}
