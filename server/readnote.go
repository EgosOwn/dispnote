package main

import (
	"fmt"
	"io"
	"time"
	"net/http"
)

var notes map[string]string

const maxNoteSize int64 = 1024
const noteIDSize int = 31
const maxNotes int = 1000000
const noteExpireSecs = time.Duration(3600) * time.Second

func getNote(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, OPTIONS")
	noteID := req.URL.Path[len("/get/"):]
	w.Header().Add("Content-Type", "text/plain")

	var note = notes[noteID]
	if note == "" {
		http.Error(w, "That note has either been read, expired, or never existed", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, notes[noteID])
	delete(notes, noteID)
}

func addNote(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "PUT, GET, OPTIONS")
	if req.Method == http.MethodOptions {
		return
	}
	if len(notes) >= maxNotes{
		http.Error(w, "Maximum number of notes on server reached", http.StatusInternalServerError)
		return
	}
	if req.Method == http.MethodPut {
		noteID := req.URL.Path[len("/put/"):]

		if len(noteID) != noteIDSize {
			http.Error(w, "Note ID size must be 31", http.StatusRequestURITooLong)
			return
		}
		if int(req.ContentLength) > int(maxNoteSize) {
			http.Error(w, "Note cannot exceed 1024 bytes", 413)
			return
		}
		if int(req.ContentLength) <= 0 {
			http.Error(w, "Note must not be empty", 413)
			return
		}

		b, _ := io.ReadAll(req.Body)

		notes[noteID] = string(b)

		go func(toDel string) {
			time.Sleep(noteExpireSecs)
			delete(notes, toDel) // Noop if key doesn't exist
		}(noteID)

	} else {
		http.Error(w, "Must PUT with body", http.StatusBadRequest)
	}
	fmt.Fprintln(w, "Note registered.")
}

func main() {

	http.HandleFunc("/get/", getNote)
	http.HandleFunc(("/put/"), addNote)

	notes = make(map[string]string)

	http.ListenAndServe(":8090", nil)
}
