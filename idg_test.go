package idg

import (
	"fmt"
	"testing"

	"github.com/missionMeteora/uuid"
)

var (
	idSink   ID
	uuidSink uuid.UUID
)

func TestIDGIndexing(t *testing.T) {
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

func testIndex(id ID, expected uint64) (err error) {
	var idx uint64
	if idx, err = id.Index(); err != nil {
		return
	} else if idx != expected {
		return fmt.Errorf("invalid index, expected %d and received %d", expected, idx)
	}

	return
}

func BenchmarkIDG_Gen(b *testing.B) {
	idg := New(0)
	for i := 0; i < b.N; i++ {
		idSink = idg.Next()
	}

	b.ReportAllocs()
}

func BenchmarkIDG_Gen_Para(b *testing.B) {
	idg := New(0)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			idSink = idg.Next()
		}
	})

	b.ReportAllocs()
}

func BenchmarkUUID_Gen(b *testing.B) {
	ug := uuid.NewGen()
	for i := 0; i < b.N; i++ {
		uuidSink = ug.New()
	}

	b.ReportAllocs()
}

func BenchmarkUUID_Gen_Para(b *testing.B) {
	ug := uuid.NewGen()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			uuidSink = ug.New()
		}
	})

	b.ReportAllocs()
}
