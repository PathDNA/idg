package idg

import (
	"os"
	"testing"

	"github.com/PathDNA/turtleDB"
)

func TestTIDGIndexing(t *testing.T) {
	var (
		db  turtleDB.DB
		id  ID
		err error
	)

	// Initialize basic funcsmap
	fm := turtleDB.FuncsMap{}
	// Initialize a new instance of tidg
	tidg := NewTIDG("test", fm)

	if db, err = turtleDB.New("tidg_test", "./test_data", fm); err != nil {
		return
	}
	defer os.RemoveAll("./test_data")
	defer db.Close()

	// Increment index four times
	if err = db.Update(func(txn turtleDB.Txn) (err error) {
		if id, err = tidg.Next(txn); err != nil {
			return
		}

		if id, err = tidg.Next(txn); err != nil {
			return
		}

		if id, err = tidg.Next(txn); err != nil {
			return
		}

		if id, err = tidg.Next(txn); err != nil {
			return
		}

		return
	}); err != nil {
		t.Fatal(err)
	}

	// Ensure our index is the proper value of 3 (4 entries, starting at an index of 0)
	if err = testIndex(id, 3); err != nil {
		t.Fatal(err)
	}

	// Increment index four times
	if err = db.Update(func(txn turtleDB.Txn) (err error) {
		if id, err = tidg.Next(txn); err != nil {
			return
		}

		if id, err = tidg.Next(txn); err != nil {
			return
		}

		if id, err = tidg.Next(txn); err != nil {
			return
		}

		if id, err = tidg.Next(txn); err != nil {
			return
		}

		return
	}); err != nil {
		t.Fatal(err)
	}

	// Ensure our index is the proper value of 7 (8 entries, starting at an index of 0)
	if err = testIndex(id, 7); err != nil {
		t.Fatal(err)
	}

	// Close database, testing for persistence
	if err = db.Close(); err != nil {
		t.Fatal(err)
	}

	// Re-initialize database
	if db, err = turtleDB.New("tidg_test", "./test_data", nil); err != nil {
		return
	}

	// Ensure our index is still proper after DB reboot
	if err = testIndex(id, 7); err != nil {
		t.Fatal(err)
	}

}
