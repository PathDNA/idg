package idg

import (
	"encoding/base64"

	"github.com/PathDNA/atoms"
	"github.com/itsmontoya/mum"
	"github.com/missionMeteora/toolkit/errors"
)

const (
	// ErrInvalidLength is returned when an ID is not 16 bytes in length
	ErrInvalidLength = errors.Error("invalid length")
)

var (
	// Base64 RawURLEncoding alias
	b64 = base64.RawURLEncoding
	// String length
	strLen = b64.EncodedLen(16)
)

// New will return a new ID generator
func New(idx uint64) (idg IDG) {
	idg.idx.Store(idx)
	return
}

// IDG is an ID generator
type IDG struct {
	mux atoms.Mux
	// Helper for binary encoding
	bw mum.BinaryWriter
	// Current index
	idx atoms.Uint64
}

// Next will return the next id
func (i *IDG) Next() (id ID) {
	// We atomically increment our current index by one.
	// It is safe to assume that our index is one less than the new value
	idx := i.idx.Add(1) - 1
	return newID(idx)
}

// Parse will parse a string id
func Parse(in string) (id ID, err error) {
	err = id.parse([]byte(in))
	return
}
