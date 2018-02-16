package main

import (
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Write([]byte("pong"))
}

func postBase64Image(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("pong"))
	w.Write([]byte("Image Received"))
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/image", postBase64Image)
	http.HandleFunc("/ping", ping)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
