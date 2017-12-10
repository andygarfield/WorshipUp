package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	os.Args = []string{"cmd", "../songjson/opensongmigrate/json"}

	http.Handle("/songlist/", songListHandler(os.Args[1]))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		if r.URL.Path == "/" || r.URL.Path == "" {
			f, _ := ioutil.ReadFile("./webapp/index.html")
			w.Write(f)
		} else {
			f, _ := ioutil.ReadFile("." + r.URL.Path)
			w.Write(f)
		}
	})
	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func songListHandler(songDir string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		songFiles, err := ioutil.ReadDir(songDir)
		if err != nil {
			panic(err)
		}

		var outJSON string
		for _, sf := range songFiles {
			b, _ := ioutil.ReadFile(songDir + "/" + sf.Name())
			outJSON += string(b)
			fmt.Println(string(b))
			break
		}

		fmt.Fprintf(w, outJSON)
	})
}
