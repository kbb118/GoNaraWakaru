package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	source := map[string]string{
		"Hello": "World",
	}

	/*
		source --- json --- gzip --- http.Response
				        \__ stdout
	*/
	gwriter := gzip.NewWriter(w)
	mulwriter := io.MultiWriter(os.Stdout, gwriter)
	json_encoder := json.NewEncoder(mulwriter)
	json_encoder.SetIndent("", "    ")
	json_encoder.Encode(source)
	// Flush を忘れない!
	gwriter.Flush()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
