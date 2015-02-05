package server

import (
	"encoding/json"
	"fmt"
	"github.com/nilbot/note.app/notes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var man = notes.NewNoteManager()

const PathPrefix = "/note/"

func RegisterHandlers() {
	r := mux.NewRouter()
	r.HandleFunc(PathPrefix, errorHandler(ListNotes)).Methods("GET")
	r.HandleFunc(PathPrefix, errorHandler(NewNote)).Methods("POST")
	r.HandleFunc(PathPrefix+"{id}", errorHandler(GetNote)).Methods("GET")
	r.HandleFunc(PathPrefix+"{id}", errorHandler(UpdateNote)).Methods("PUT")
	r.HandleFunc(PathPrefix+"#{tag}", errorHandler(Filter)).Methods("GET")
	http.Handle(PathPrefix, r)
}

// badRequest is handled by setting the status code in the reply to StatusBadRequest.
type badRequest struct{ error }

// notFound is handled by setting the status code in the reply to StatusNotFound.
type notFound struct{ error }

// errorHandler wraps a function returning an error by handling the error and returning a http.Handler.
// If the error is of the one of the types defined above, it is handled as described for every type.
// If the error is of another type, it is considered as an internal error and its message is logged.
func errorHandler(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err == nil {
			return
		}
		switch err.(type) {
		case badRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case notFound:
			http.Error(w, "note not found", http.StatusNotFound)
		default:
			log.Println(err)
			http.Error(w, "oops", http.StatusInternalServerError)
		}
	}
}

// ListNotes handles GET requests on /note.
// There's no parameters and it returns an object with a Notes field containing a list of tasks.
//
// Example:
//
//   req: GET /note/
//   res: 200 {"Notes": [
//          {"ID": ff5dsasd2, "Content": "Write a Note with #tag"},
//          {"ID": abcdedfg1, "Content": "Buy bread #todo"}
//          ],
//          "Tags": ["tag"]
//          }
func ListNotes(w http.ResponseWriter, r *http.Request) error {
	res := struct {
		Notes []*notes.Note
		Tags  []string
	}{
		man.AllNotes(),
		man.AllTags(),
	}
	return json.NewEncoder(w).Encode(res)
}

// NewNote handles POST requests on /note.
// The request body must contain a JSON object with a Content field.
// The status code of the response is used to indicate any error.
//
// Examples:
//
//   req: POST /note/ {"Content": ""}
//   res: 400 empty content
//
//   req: POST /note/ {"Content": "Buy milk"}
//   res: 200
func NewNote(w http.ResponseWriter, r *http.Request) error {
	req := struct{ Content string }{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return badRequest{err}
	}
	t, err := notes.NewNote(req.Content)
	if err != nil {
		return badRequest{err}
	}
	return man.Save(t)
}

// parseID obtains the id variable from the given request url,
// parses the obtained text and returns the result.
func parseID(r *http.Request) (string, error) {
	txt, ok := mux.Vars(r)["id"]
	if !ok {
		return "", fmt.Errorf("note id not found")
	}
	return txt, nil
}

// parseHashtag obtains the tag variable from the given request url,
// parses the obtained text and returns the result.
func parseHashtag(r *http.Request) (string, error) {
	txt, ok := mux.Vars(r)["tag"]
	if !ok {
		return "", fmt.Errorf("tag not found")
	}
	return txt, nil
}

// GetNote handles GET requsts to /note/{ID}.
// There's no parameters and it returns a JSON encoded note.
//
// Examples:
//
//   req: GET /note/abcdefg123
//   res: 200 {"ID": abcdefg123, "Content": "Buy milk"}
//
//   req: GET /note/4242424242
//   res: 404 note not found
func GetNote(w http.ResponseWriter, r *http.Request) error {
	id, err := parseID(r)
	log.Println("Note is ", id)
	if err != nil {
		return badRequest{err}
	}
	n, ok := man.Find(id)
	log.Println("Found", ok)

	if !ok {
		return notFound{}
	}
	return json.NewEncoder(w).Encode(n)
}

// UpdateNote handles PUT requests to /note/{ID}.
// The request body must contain a JSON encoded note.
//
// Example:
//
//   req: PUT /note/1234abcd {"ID": 1234abcd, "Content": "rewrite note"}
//   res: 200
//
//   req: PUT /note/42 {"ID": 42, "Content": "Write anything"}
//   res: 400 inconsistent note IDs
func UpdateNote(w http.ResponseWriter, r *http.Request) error {
	id, err := parseID(r)
	if err != nil {
		return badRequest{err}
	}
	var n notes.Note
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		return badRequest{err}
	}
	if n.ID != id {
		return badRequest{fmt.Errorf("inconsistent note IDs")}
	}
	if _, ok := man.Find(id); !ok {
		return notFound{}
	}
	return man.Save(&n)
}

// Filter with tag handles GET requests to /note/#{Tag}.
func Filter(w http.ResponseWriter, r *http.Request) error {
	tag, err := parseHashtag(r)
	log.Println("Tag is", tag)
	if err != nil {
		return badRequest{err}
	}
	ns, ok := man.NotesWith(tag)
	log.Println("Found", ok)
	if !ok {
		return notFound{}
	}
	return json.NewEncoder(w).Encode(ns)
}
