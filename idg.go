package idg

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/PathDNA/atoms"
	"github.com/itsmontoya/mum"
	"github.com/missionMeteora/toolkit/errors"
)

const (
	// ErrInvalidLength is returned when an ID is not 16 bytes in length
	ErrInvalidLength = errors.Error("invalid length")
)

var (
	// Base64 encoder alias
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

func newID(idx uint64) (id ID) {
	// Current Unix timestamp (in nanoseconds)
	now := time.Now().UnixNano()
	// Helper for binary encoding
	var bw mum.BinaryWriter
	// Copy index bytes to first 8 bytes
	copy(id[:8], bw.Uint64(idx))
	// Copy unix timestamp bytes to last 8 bytes
	copy(id[8:], bw.Int64(now))
	return
}

// ID represents an id
type ID [16]byte

func (id *ID) parse(in []byte) (err error) {
	if len(in) != strLen {
		// Decoded value has to be 16 bytes or it's not valid
		err = ErrInvalidLength
		return
	}

	// Decode inbound bytes as base64
	// Write the bytes directly to our array
	_, err = b64.Decode(id[:], in)
	return
}

// Index will return the index of an ID
func (id *ID) Index() (idx uint64, err error) {
	// Helper for binary decoding
	var br mum.BinaryReader
	// Grab the index from the first 8 bytes
	return br.Uint64((*id)[:8])
}

// Time will return the time.Time of an ID
func (id *ID) Time() (t time.Time, err error) {
	var (
		// Helper for binary decoding
		br mum.BinaryReader
		// Timestamp
		ts int64
	)

	// Grab the Unix timestamp from the last 8 bytes
	if ts, err = br.Int64((*id)[8:]); err != nil {
		return
	}

	// Parse Unix timestamp (as nanoseconds)
	t = time.Unix(0, ts)
	return
}

// Bytes will return the byteslice representation
// Note: This function is unsafe and can change the underlying array
// Please.. read only!
func (id *ID) Bytes() (out []byte) {
	out = (*id)[:]
	return
}

// String will return a string representation
// Note: This is referenced as a non-pointer so it can be called directly
// from a struct utilizing the non-pointer value of ID
func (id *ID) String() (out string) {
	out = b64.EncodeToString(id[:])
	return
}

// IsEmpty will return if an ID is empty
func (id *ID) IsEmpty() (empty bool) {
	for i := 0; i < 16; i++ {
		if (*id)[i] != 0 {
			return
		}
	}
	return true
}

// MarshalJSON is a JSON encoding helper func
func (id ID) MarshalJSON() (out []byte, err error) {
	return json.Marshal(id.String())
}

// UnmarshalJSON is a JSON decoding helper func
func (id ID) UnmarshalJSON(in []byte) (err error) {
	var str string
	if err = json.Unmarshal(in, &str); err != nil {
		return
	}

	stripped := in[1 : len(in)-1]
	return id.parse(stripped)
}

// Parse will parse a string id
func Parse(in string) (id ID, err error) {
	err = id.parse([]byte(in))
	return
}
