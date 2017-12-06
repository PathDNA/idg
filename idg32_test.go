package idg

import (
	"fmt"
	"testing"
)

var (
	id32Sink ID32
)

func TestIDG32Indexing(t *testing.T) {
	var err error
	// Generate new ID with an index starting at 3
	idg := New32(3)
	if err = testIndex32(idg.Next(), 3); err != nil {
		t.Fatal(err)
	}

	if err = testIndex32(idg.Next(), 4); err != nil {
		t.Fatal(err)
	}

	if err = testIndex32(idg.Next(), 5); err != nil {
		t.Fatal(err)
	}

	id := idg.Next()
	fmt.Println(id.String())
}

func testIndex32(id ID32, expected uint32) (err error) {
	var idx uint32
	if idx, err = id.Index(); err != nil {
		return
	} else if idx != expected {
		return fmt.Errorf("invalid index, expected %d and received %d", expected, idx)
	}

	return
}

func BenchmarkIDG32_Gen(b *testing.B) {
	idg := New32(0)
	for i := 0; i < b.N; i++ {
		id32Sink = idg.Next()
	}

	b.ReportAllocs()
}

func BenchmarkIDG32_Gen_Para(b *testing.B) {
	idg := New32(0)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			id32Sink = idg.Next()
		}
	})

	b.ReportAllocs()
}
