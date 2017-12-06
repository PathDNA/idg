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
	// Empty ID used for matching
	emptyID   = ID{}
	emptyID32 = ID32{}
)

// New will return a new ID generator
func New(idx uint64) (idg IDG) {
	idg.idx.Store(idx)
	return
}

// IDG is an non-persistent atomic ID generator
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
	return newID(idx, -1)
}

// Next32 will return the next 32-bit id
func (i *IDG) Next32() (id ID32) {
	// We atomically increment our current index by one.
	// It is safe to assume that our index is one less than the new value
	idx := i.idx.Add(1) - 1
	return newID32(uint32(idx), -1)
}
