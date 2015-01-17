package notes

import (
	"crypto/sha512"
	"fmt"
	"io"
	"testing"
)

func newNoteOrFatal(t *testing.T, content string) *Note {
	note, err := NewNote(content)
	if err != nil {
		t.Fatalf("new note: %v", err)
	}
	return note
}

func TestNewNote(t *testing.T) {
	content := "copied from campoy/todo"
	hash := sha512.New()
	io.WriteString(hash, content)
	sha512_checksum := fmt.Sprintf("%x", hash.Sum(nil))
	note := newNoteOrFatal(t, content)
	if note.Content != content {
		t.Errorf("expected content %q, got %q", content, note.Content)
	}
	if note.ID != sha512_checksum {
		t.Errorf("expected checksum %q, got %q", sha512_checksum, note.ID)
	}

}

func TestSave(t *testing.T) {
	content := "test a new note with a single hashtag #test, without linebreaks etc."
	note := newNoteOrFatal(t, content)
	man := NewNoteManager()
	man.Save(note)
	all := man.AllNotes()
	if len(all) != 1 {
		t.Errorf("expected 1 note, got %v", len(all))
	}
	if *all[0] != *note {
		t.Errorf("expected %v, got %v", note, all[0])
	}
}

func TestParse(t *testing.T) {
	content := "test a new note with a single hashtag #test, without linebreaks etc."
	tags := parse(content)
	if len(tags) != 1 {
		t.Errorf("expected 1 tag, got %v", len(tags))
	}
	if tags[0] != "test" {
		t.Errorf("expected parsed tag %q, got %q", "test", tags[0])
	}
}

func getNoteWithTagAndSaveInManager(t *testing.T) (*NoteManager, *Note, string, []string) {
	content := "test a new note with a single hashtag #test, without linebreaks etc."
	tags := parse(content)
	note := newNoteOrFatal(t, content)
	man := NewNoteManager()
	man.Save(note)
	return man, note, content, tags
}

func TestSaveWithVerifyingTag(t *testing.T) {
	content := "test a new note with a single hashtag #test, without linebreaks etc."
	tag := "test"
	note := newNoteOrFatal(t, content)
	man := NewNoteManager()
	man.Save(note)
	tags := man.AllTags()
	if len(tags) != 1 {
		t.Errorf("expected 1 tag, got %v", len(tags))
	}
	if _, ok := man.NotesWith(tag); !ok {
		t.Errorf("expected to find tag %q, got nothing", "test")
	}
	if val, ok := man.NotesWith(tag); ok {
		for _, s := range val {
			if s != note.ID {
				t.Errorf("expected to find mapping of %q to %q, got %q to %q", "test", note.ID, "test", val)
			}
		}
	}
}
