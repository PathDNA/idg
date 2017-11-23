package idg

import (
	"encoding/json"
	"testing"
	"time"
)

func TestIDTime(t *testing.T) {
	var (
		tt  time.Time
		err error
	)

	// Get a current timestamp
	now := time.Now()
	// Ensure our new ID is at least one millisecond behind our timestamp
	time.Sleep(time.Millisecond)
	// Generate a new ID
	id := newID(0, -1)
	// Get the time from our ID
	if tt, err = id.Time(); err != nil {
		t.Fatal(err)
	}
	// Check to see if our id's time is after our timestamp (it should be)
	if !tt.After(now) {
		t.Fatalf("invalid time, should be after initial timestamp: %v / %v", now, tt)
	}
}

func TestIDParse(t *testing.T) {
	var (
		nid ID
		err error
	)

	// Generate an ID with the index starting at 1337
	id := newID(1337, -1)
	// Get the string representation of our ID
	sid := id.String()
	// Parse the string to a new ID
	if nid, err = Parse(sid); err != nil {
		t.Fatal(err)
	}
	// Check if the ID's match
	if id != nid {
		t.Fatalf("ID's do not match: %v / %v", id.Bytes(), nid.Bytes())
	}
}

func TestJSON(t *testing.T) {
	var (
		b   []byte
		err error
	)

	// Generate an ID with the index starting at 1337
	id := newID(1337, -1)
	// Marshal ID as JSON
	if b, err = json.Marshal(&id); err != nil {
		t.Fatal(err)
	}

	var nid ID
	// Parse as JSON to a new ID
	if err = json.Unmarshal(b, &nid); err != nil {
		t.Fatal(err)
	}
	// Check if the ID's match
	if id != nid {
		t.Fatalf("ID's do not match: %v / %v", id.Bytes(), nid.Bytes())
	}
}
