# bench-json

test the memory usage of 1000 tables during encoding and decoding.

each table takes about 7 kb, including 10 columns and 10 indexes.
```
 go test -bench=BenchmarkJsonForUnmarshalParallel -benchtime=10s -memprofile=memPar.prof
```

1  goroutine,  decoding takes about 1.2 GB

5  goroutines, decoding takes about 2.0 GB

10 goroutines, decoding takes about 2.43GB

we see the encoding and decoding process costs too much memory.

# command
```
 go test -bench=BenchmarkJsonForUnmarshalParallel -benchtime=5s -memprofile=memPar.prof
 go tool pprof -http=:8080 memPar.prof 
```

# GC
```
 go test -bench=BenchmarkJsonForUnmarshalParallel -benchtime=5s -memprofile=memPar.prof
```
after invoke `runtime.GC()`after each json.unmarshal(), the memory usage drops so much.

1  goroutine,  691 MB -> 90 MB

10 goroutine, 3770 MB -> 355 MB

# json vs. protobuf
```
mac:json-bench fzh$  go test -bench=.* -benchmem -count=2
goos: darwin
goarch: amd64
pkg: github.com/json-bench
BenchmarkJsonForUnmarshalNoGC-12                       2         767289374 ns/op        83426280 B/op    1390010 allocs/op
BenchmarkJsonForUnmarshalNoGC-12                       2         742245306 ns/op        83426280 B/op    1390010 allocs/op
BenchmarkProtoForUnmarshalNoGC-12                     21          53979592 ns/op        74626314 B/op     960020 allocs/op
BenchmarkProtoForUnmarshalNoGC-12                     20          60220235 ns/op        74626320 B/op     960020 allocs/op
BenchmarkJsonForUnmarshalParallelNoGC-12               7         160865067 ns/op        83206891 B/op    1390110 allocs/op
BenchmarkJsonForUnmarshalParallelNoGC-12               7         164963511 ns/op        83206093 B/op    1390106 allocs/op
BenchmarkProteForUnmarshalParallelNoGC-12             80          18263411 ns/op        74404418 B/op     960113 allocs/op
BenchmarkProteForUnmarshalParallelNoGC-12             81          18263492 ns/op        74404402 B/op     960113 allocs/op
```
The protobuf is almost 10X faster than json in serialization.
in terms of memory usage, protobuf also takes less memory.

TODO: how to implement `interface{}` in protobuf?
