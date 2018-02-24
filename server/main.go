package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func postBase64Image(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("pong"))
	r.ParseForm()
	var data = r.Form["frame_0"][0]
	// // The actual image starts after the ","
	i := strings.Index(data, ",")
	if i < 0 {
		log.Fatal("no comma")
	}
	dec, err := base64.StdEncoding.DecodeString(data[i+1:])
	if err != nil {
		panic(err)
	}

	// os.MkdirAll("/output/", os.FileMode(0755))

	f, err := os.Create("output/frame_0.png")
	defer f.Close()
	if _, err := f.Write(dec); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}
	w.Write([]byte("Image Received"))
	// Return image in response
	// pass reader to NewDecoder
	// dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[i+1:]))
	// w.Header().Set("Content-Type", "image/png")
	// io.Copy(w, dec)
}

func main() {
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir(".")))
	r.HandleFunc("/image", postBase64Image)
	r.HandleFunc("/ping", ping)
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
