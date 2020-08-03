# bench-json

test the memory usage of 1000 tables during encoding and decoding.

1 goroutine, 1.2 GB
5 goroutine, 2.0 GB

# command
 go test -bench=BenchmarkJsonForUnmarshalParallel -benchtime=5s -memprofile=memPar.prof
 go tool pprof -http=:8080 mem.prof 

