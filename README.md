# idg (ID generator)

# Benchmarks
```bash
## idg
BenchmarkGenerationIDG-16             20000000    96.4 ns/op    0 B/op    0 allocs/op
BenchmarkGenerationParallelIDG-16     30000000    36.6 ns/op    0 B/op    0 allocs/op
## missionMeteora/uuid
BenchmarkGenerationUUID-16            10000000    120 ns/op     0 B/op    0 allocs/op
BenchmarkGenerationParallelUUID-16    5000000     291 ns/op     0 B/op    0 allocs/op
```