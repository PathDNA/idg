package main

import (
	"fmt"
	"time"

	"github.com/PathDNA/idg"
)

func main() {
	var (
		idx uint64
		t   time.Time
		err error
	)

	// Initialize id generator
	idg := idg.New(uint64(1337))
	// Get next ID
	id := idg.Next()
	// Get string representation of ID
	str := id.String()
	// Get index of ID
	if idx, err = id.Index(); err != nil {
		panic("Error getting index: " + err.Error())
	}
	// Get time of ID
	if t, err = id.Time(); err != nil {
		panic("Error getting time: " + err.Error())
	}
	// Display ID information
	fmt.Printf("ID\nString: %s\nIndex: %d\nTime: %v\n", str, idx, t)
}
