# bench-json

test the memory usage of 1000 tables during encoding and decoding.

each table takes about 7 kb, including 10 columns and 10 indexes.

1  goroutine,  decoding takes about 1.2 GB

5  goroutines, decoding takes about 2.0 GB

10 goroutines, decoding takes about 2.43GB

we see the encoding and decoding process costs too much memory.

# command
```
 go test -bench=BenchmarkJsonForUnmarshalParallel -benchtime=5s -memprofile=memPar.prof
 go tool pprof -http=:8080 mem.prof 
```
