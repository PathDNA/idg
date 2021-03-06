package idg

import (
	"encoding/json"
	"time"

	"github.com/itsmontoya/mum"
)

// newID32 will return a new ID with the provided index and timestamp
// Note: If timestamp is set to -1, the current Unix timestamp will
// be utilized
func newID32(idx uint32, ts int64) (id ID32) {
	// Helper for binary encoding
	var bw mum.BinaryWriter
	// Check if timestamp is set (or needs to be set)
	if ts == -1 {
		// Timestamp is set to -1, set timestamp to current Unix timestamp (in seconds)
		// Note: Seconds was decided to be utilized instead of nanoseconds
		// To aid in an easier integration with Javascript for front-end clients
		// utilizing idg. Technically, we could utilize milliseconds and maintain
		// Javascript compatibility. That being said, seconds feels like a much
		// more universal Unix time reference interval.
		ts = time.Now().Unix()
	}
	// Copy index bytes to first 8 bytes
	copy(id[:4], bw.Uint32(uint32(idx)))
	// Copy unix timestamp bytes to last 8 bytes
	copy(id[4:], bw.Uint32(uint32(ts)))
	return
}

// ID32 represents a 32-bit id
type ID32 [8]byte

func (id *ID32) parse(in []byte) (err error) {
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
func (id *ID32) Index() (idx uint32, err error) {
	// Helper for binary decoding
	var br mum.BinaryReader
	// Check if ID is nil
	if id == nil {
		// ID is nil, return early
		err = ErrEmptyID
		return
	}
	// Grab the index from the first 8 bytes
	return br.Uint32(id[:4])
}

// Time will return the time.Time of an ID
func (id *ID32) Time() (t time.Time, err error) {
	var (
		// Helper for binary decoding
		br mum.BinaryReader
		// Timestamp
		ts uint32
	)
	// Check if ID is nil
	if id == nil {
		// ID is nil, return early
		err = ErrEmptyID
		return
	}
	// Grab the Unix timestamp from the last 8 bytes
	if ts, err = br.Uint32(id[4:]); err != nil {
		return
	}

	// Parse Unix timestamp (as nanoseconds)
	t = time.Unix(int64(ts), 0)
	return
}

// Bytes will return the byteslice representation
// Note: This function is unsafe and can change the underlying array
// Please.. read only!
func (id *ID32) Bytes() (out []byte) {
	if id == nil {
		return
	}

	out = id[:]
	return
}

// String will return a string representation
// Note: This is referenced as a non-pointer so it can be called directly
// from a struct utilizing the non-pointer value of ID
func (id *ID32) String() (out string) {
	if id == nil {
		return
	}

	out = b64.EncodeToString(id[:])
	return
}

// IsEmpty will return if an ID is empty
func (id *ID32) IsEmpty() (empty bool) {
	return id == nil || *id == emptyID32
}

// MarshalJSON is a JSON encoding helper func
func (id *ID32) MarshalJSON() (out []byte, err error) {
	// Check if ID is nil
	if id == nil {
		return
	}

	return json.Marshal(id.String())
}

// UnmarshalJSON is a JSON decoding helper func
func (id *ID32) UnmarshalJSON(in []byte) (err error) {
	var str string
	// Unmarshal inbound value as a string
	if err = json.Unmarshal(in, &str); err != nil {
		return
	}
	// Strip double-quotation from head and tail
	stripped := in[1 : len(in)-1]
	// Return result of the parsed value
	return id.parse(stripped)
}
