package notes

import (
	"crypto/sha512"
	"fmt"
	"io"
	"strings"
)

type Note struct {
	ID      string
	Content string
}

type NoteManager struct {
	tags  map[string][]string
	notes map[string]*Note
}

func NewNoteManager() *NoteManager {
	return &NoteManager{
		tags:  make(map[string][]string),
		notes: make(map[string]*Note),
	}
}
func NewNote(content string) (*Note, error) {
	if content == "" {
		return nil, fmt.Errorf("empty content")
	}

	hash := sha512.New()
	io.WriteString(hash, content)

	result := &Note{
		fmt.Sprintf("%x", hash.Sum(nil)),
		content,
	}
	return result, nil
}

func parse(str string) []string {
	if strings.ContainsRune(str, '#') {
		str = strings.TrimSpace(str)
		str = strings.Replace(str, ",", " ", -1)
		str = strings.Replace(str, ".", " ", -1)
		str = strings.Replace(str, ":", " ", -1)
		str = strings.Replace(str, "\"", " ", -1)
		str = strings.Replace(str, ";", " ", -1)
		words := strings.Split(str, " ")
		var a []string
		for _, v := range words {
			if strings.Index(v, "#") == 0 {
				a = append(a, v[1:len(v)])
			}
		}
		return a
	}
	return nil
}

func (man *NoteManager) Save(n *Note) error {
	man.notes[n.ID] = n
	t := parse(n.Content)
	for _, v := range t {
		man.tags[v] = append(man.tags[v], n.ID)
	}
	return nil
}

func (man *NoteManager) AllNotes() []*Note {
	v := make([]*Note, len(man.notes))
	idx := 0
	for _, value := range man.notes {
		v[idx] = value
		idx++
	}
	return v
}

func (man *NoteManager) AllTags() []string {
	v := make([]string, len(man.tags))
	idx := 0
	for key, _ := range man.tags {
		v[idx] = key
		idx++
	}
	return v
}

func (man *NoteManager) NotesWith(tag string) ([]string, bool) {
	value, ok := man.tags[tag]
	return value, ok
}
