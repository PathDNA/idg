# idg (ID generator)
idg is a simple ID generation helper. Each ID is indexed and assigned a timestamp upon creation.

# Usage
```go
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
	idg := idg.New(0)
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

```


# Benchmarks
```bash
## idg
BenchmarkGenerationIDG-16             20000000    96.4 ns/op    0 B/op    0 allocs/op
BenchmarkGenerationParallelIDG-16     30000000    36.6 ns/op    0 B/op    0 allocs/op

## missionMeteora/uuid
BenchmarkGenerationUUID-16            10000000    120 ns/op     0 B/op    0 allocs/op
BenchmarkGenerationParallelUUID-16    5000000     291 ns/op     0 B/op    0 allocs/op
```