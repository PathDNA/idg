package idg

import (
	"os"
	"testing"
)

func TestPIDGIndexing(t *testing.T) {
	var (
		pidg *PIDG
		id   ID
		err  error
	)
	defer os.RemoveAll("./test_data")

	// Initialize a persistent id generator
	if pidg, err = NewPersistent("test1", "./test_data"); err != nil {
		t.Fatal(err)
	}
	defer pidg.Close()

	if id, err = pidg.Next(); err != nil {
		t.Fatal(err)
	}

	if err = testIndex(id, 0); err != nil {
		t.Fatal(err)
	}

	if id, err = pidg.Next(); err != nil {
		t.Fatal(err)
	}

	if err = testIndex(id, 1); err != nil {
		t.Fatal(err)
	}

	if id, err = pidg.Next(); err != nil {
		t.Fatal(err)
	}

	if err = testIndex(id, 2); err != nil {
		t.Fatal(err)
	}

	if err = pidg.Close(); err != nil {
		t.Fatal(err)
	}

	// Re-initialize a persistent id generator to test data loading
	if pidg, err = NewPersistent("test1", "./test_data"); err != nil {
		t.Fatal(err)
	}

	if id, err = pidg.Next(); err != nil {
		t.Fatal(err)
	}
	// Value should now be 3
	if err = testIndex(id, 3); err != nil {
		t.Fatal(err)
	}

	var id32 ID32
	if id32, err = pidg.Next32(); err != nil {
		t.Fatal(err)
	}

	if err = testIndex32(id32, 4); err != nil {
		t.Fatal(err)
	}

}

func BenchmarkPIDG_Gen(b *testing.B) {
	var (
		pidg *PIDG
		err  error
	)
	defer os.RemoveAll("./test_data")

	// Generate new ID with an index starting at 3
	if pidg, err = NewPersistent("test1", "./test_data"); err != nil {
		// Error encountered, bail out
		b.Fatal(err)
	}
	defer pidg.Close()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if idSink, err = pidg.Next(); err != nil {
			// Error encountered, bail out
			b.Fatal(err)
		}
	}

	b.ReportAllocs()
}

func BenchmarkPIDG_Gen_Para(b *testing.B) {
	var (
		pidg *PIDG
		err  error
	)
	defer os.RemoveAll("./test_data")

	// Generate new ID with an index starting at 3
	if pidg, err = NewPersistent("test1", "./test_data"); err != nil {
		// Error encountered, bail out
		b.Fatal(err)
	}
	defer pidg.Close()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		var err error
		for pb.Next() {
			if idSink, err = pidg.Next(); err != nil {
				b.Fatal()
			}
		}
	})

	b.ReportAllocs()
}
