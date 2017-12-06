package idg

import (
	"io"
	"os"
	"path"

	"github.com/PathDNA/atoms"
	"github.com/itsmontoya/mum"
)

//NewPersistent will return a new ID generator
func NewPersistent(key, dir string) (pidg *PIDG, err error) {
	var p PIDG
	// Set file
	if err = p.setFile(key, dir); err != nil {
		return
	}
	// Set encoder
	p.enc = mum.NewEncoder(p.pf)
	pidg = &p
	return
}

// PIDG is an non-persistent atomic ID generator
type PIDG struct {
	mux atoms.Mux
	// Helper for binary encoding
	bw mum.BinaryWriter
	// Persistance file
	pf *os.File
	// Encoder writer
	enc *mum.Encoder
	// Current index
	idx uint64
}

// setFile will set the internal persistence file
func (p *PIDG) setFile(key, dir string) (err error) {
	// Ensure all directories exist
	if err = os.MkdirAll(dir, 0744); err != nil {
		return
	}

	// Filepath
	fp := path.Join(dir, key+".idg")
	// Open (or create) file with read/write functionality
	if p.pf, err = os.OpenFile(fp, os.O_CREATE|os.O_RDWR, 0644); err != nil {
		return
	}
	// Create a temporary decoder to read initial value
	dec := mum.NewDecoder(p.pf)
	var idx uint64
	// Decode a uint64 value from the file
	if idx, err = dec.Uint64(); err != nil {
		if err == io.EOF {
			// io.EOF means we do not yet have any index data saved, our default index of 0 is
			// our intended value, error does not need to be reported
			err = nil
		}

		return
	}
	// Current index would be the NEXT index following the last persisted value
	p.idx = idx + 1
	return
}

// persist will store the current index to disk
// Note: This is NOT thread-safe, please ensure locking is
// handled by the calling func
func (p *PIDG) persist() (err error) {
	if _, err = p.pf.Seek(0, io.SeekStart); err != nil {
		return
	}

	return p.enc.Uint64(p.idx)
}

// Next will return the next id
func (p *PIDG) Next() (id ID, err error) {
	var idx uint64
	p.mux.Update(func() {
		idx = p.idx
		// Perist value to disk
		if err = p.persist(); err != nil {
			return
		}
		// Increment index value
		p.idx++
	})
	// Break early if error exists
	if err != nil {
		return
	}
	// Set id with the retrieved index (utilizing a current timestamp)
	id = newID(idx, -1)
	return
}

// Next32 will return the next 32-bit id
func (p *PIDG) Next32() (id ID32, err error) {
	var idx uint64
	p.mux.Update(func() {
		idx = p.idx
		// Perist value to disk
		if err = p.persist(); err != nil {
			return
		}
		// Increment index value
		p.idx++
	})
	// Break early if error exists
	if err != nil {
		return
	}
	// Set id with the retrieved index (utilizing a current timestamp)
	id = newID32(uint32(idx), -1)
	return
}

// Close will close the internal file
func (p *PIDG) Close() (err error) {
	p.mux.Update(func() {
		err = p.pf.Close()
	})

	return
}
