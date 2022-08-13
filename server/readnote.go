package main

import (
	"fmt"
	"io"
	"time"
	"net/http"
)

var notes map[string]string

const maxNoteSize int64 = 1024
const noteIDSize int = 22
const maxNotes int = 1000000
const noteExpireSecs = time.Duration(3600) * time.Second

func getNote(w http.ResponseWriter, req *http.Request) {
	noteID := req.URL.Path[len("/getnote/"):]
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	var note = notes[noteID]
	if note == "" {
		fmt.Fprintf(w, "That note has either been read, expired, or never existed")
		return
	}
	fmt.Fprintf(w, notes[noteID])
	delete(notes, noteID)
}

func addNote(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	if (len(notes) >= maxNotes){
		http.Error(w, "Maximum number of notes on server reached", http.StatusInternalServerError)
		return
	}
	if req.Method == "PUT" {
		noteID := req.URL.Path[len("/put/"):]

		if len(noteID) != noteIDSize {
			fmt.Fprintf(w, "Note ID size must be %d", noteIDSize)
			http.Error(w, "", http.StatusRequestURITooLong)
			return
		}
		if int(req.ContentLength) > int(maxNoteSize) {
			fmt.Fprintf(w, "Note cannot exceed %d bytes", noteIDSize)
			http.Error(w, "", 413)
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
			delete(notes, toDel)
		}(noteID)

	} else {
		http.Error(w, "Must PUT with body", http.StatusBadRequest)
	}
	fmt.Fprintln(w, "Note registered.")
}

func main() {

	http.HandleFunc("/getnote/", getNote)
	http.HandleFunc(("/put/"), addNote)

	notes = make(map[string]string)

	http.ListenAndServe(":8090", nil)
}
