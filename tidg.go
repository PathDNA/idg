package idg

import (
	"encoding/json"

	"github.com/PathDNA/turtleDB"
	"github.com/itsmontoya/mum"
)

const tidgBkt = "__idg"

// NewTIDG will return a new turtleDB-backed ID generator
func NewTIDG(key string, fm turtleDB.FuncsMap) (t TIDG) {
	t.key = key
	fm.Put(tidgBkt, marshalIndex, unmarshalIndex)
	return
}

// TIDG is a persistent turtleDB-based ID generator
type TIDG struct {
	// Helper for binary encoding
	bw mum.BinaryWriter
	// Key utilized for the index value
	key string
}

func (t *TIDG) getIndex(bkt turtleDB.Bucket) (idx uint64, err error) {
	var val turtleDB.Value
	// Get value set for our TIDG.key
	if val, err = bkt.Get(t.key); err != nil {
		if err == turtleDB.ErrKeyDoesNotExist {
			// If the key does not exist, we can set error to nil
			// An index of 0 will be just as intended
			err = nil
		}

		return
	}

	var ok bool
	if idx, ok = val.(uint64); !ok {
		// Index is not the right type, abort!
		err = turtleDB.ErrInvalidType
		return
	}

	return
}

// Next will return the next id
func (t *TIDG) Next(txn turtleDB.Txn) (id ID, err error) {
	var bkt turtleDB.Bucket
	// Ensure idg bucket exists
	if bkt, err = txn.Create("__idg"); err != nil {
		// Error encountered while creating idg bucket
		return
	}

	var idx uint64
	// Get current index
	if idx, err = t.getIndex(bkt); err != nil {
		// We encountered an error while getting, return
		return
	}

	// Increment index value and set it as the index for our TIDG.key
	if err = bkt.Put(t.key, idx+1); err != nil {
		// We encountered an error while putting, return
		return
	}

	id = newID(idx, -1)
	return
}

// marshalIndex is an encoding helper function for turtleDB
func marshalIndex(val turtleDB.Value) (b []byte, err error) {
	var (
		idx uint64
		ok  bool
	)

	if idx, ok = val.(uint64); !ok {
		err = turtleDB.ErrInvalidType
		return
	}

	return json.Marshal(idx)
}

// unmarshalIndex is an decoding helper function for turtleDB
func unmarshalIndex(b []byte) (val turtleDB.Value, err error) {
	var idx uint64
	if err = json.Unmarshal(b, &idx); err != nil {
		return
	}

	val = idx
	return
}
