package idg

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/missionMeteora/uuid"
)

var (
	idSink   ID
	uuidSink uuid.UUID
)

func TestTime(t *testing.T) {
	var (
		tt  time.Time
		err error
	)

	// Get a current timestamp
	now := time.Now()
	// Ensure our new ID is at least one millisecond behind our timestamp
	time.Sleep(time.Millisecond)
	// Generate a new ID
	id := newID(0)
	// Get the time from our ID
	if tt, err = id.Time(); err != nil {
		t.Fatal(err)
	}
	// Check to see if our id's time is after our timestamp (it should be)
	if !tt.After(now) {
		t.Fatalf("invalid time, should be after initial timestamp: %v / %v", now, tt)
	}
}
func TestIndexing(t *testing.T) {
	var err error
	// Generate new ID with an index starting at 3
	idg := New(3)
	if err = testIndex(idg.Next(), 3); err != nil {
		t.Fatal(err)
	}

	if err = testIndex(idg.Next(), 4); err != nil {
		t.Fatal(err)
	}

	if err = testIndex(idg.Next(), 5); err != nil {
		t.Fatal(err)
	}
}

func TestParse(t *testing.T) {
	var (
		nid ID
		err error
	)

	// Generate an ID with the index starting at 1337
	id := newID(1337)
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
	id := newID(1337)
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

func testIndex(id ID, expected uint64) (err error) {
	var idx uint64
	if idx, err = id.Index(); err != nil {
		return
	} else if idx != expected {
		return fmt.Errorf("invalid index, expected %d and received %d", expected, idx)
	}

	return
}

func BenchmarkGenerationIDG(b *testing.B) {
	idg := New(0)
	for i := 0; i < b.N; i++ {
		idSink = idg.Next()
	}

	b.ReportAllocs()
}

func BenchmarkGenerationUUID(b *testing.B) {
	ug := uuid.NewGen()
	for i := 0; i < b.N; i++ {
		uuidSink = ug.New()
	}

	b.ReportAllocs()
}

func BenchmarkGenerationParallelIDG(b *testing.B) {
	idg := New(0)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			idSink = idg.Next()
		}
	})

	b.ReportAllocs()
}

func BenchmarkGenerationParallelUUID(b *testing.B) {
	ug := uuid.NewGen()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			uuidSink = ug.New()
		}
	})

	b.ReportAllocs()
}
