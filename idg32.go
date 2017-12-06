package idg

import (
	"github.com/PathDNA/atoms"
	"github.com/itsmontoya/mum"
)

// New32 will return a new 32-bit ID generator
func New32(idx uint32) (idg IDG32) {
	idg.idx.Store(idx)
	return
}

// IDG32 is an non-persistent atomic ID generator
type IDG32 struct {
	mux atoms.Mux
	// Helper for binary encoding
	bw mum.BinaryWriter
	// Current index
	idx atoms.Uint32
}

// Next will return the next id
func (i *IDG32) Next() (id ID32) {
	// We atomically increment our current index by one.
	// It is safe to assume that our index is one less than the new value
	idx := i.idx.Add(1) - 1
	return newID32(idx, -1)
}
