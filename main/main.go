package main

import (
	"fmt"
	_ "github.com/gogo/protobuf/proto"
	json_bench "github.com/json-bench"
	"runtime"
	"sync/atomic"
	"time"
)
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func PrintMemUsage(stop *int32) {
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	var m runtime.MemStats
	var last uint64
	last = 0
	for(atomic.LoadInt32(stop)==1) {
		runtime.ReadMemStats(&m)
		cur := bToMb(m.TotalAlloc)
		if cur - last > 10 || last - cur >10{
			last = cur
			fmt.Println("instance mem: ",  "Alloc ",bToMb(m.Alloc), "TotalAlloc ",bToMb(m.TotalAlloc),"Sys ",bToMb(m.Sys),"NumGC ",m.NumGC)
		}
		time.Sleep(500* time.Nanosecond)
	}
}


func json() {
	tblByte := json_bench.JsonFor()
	var stop int32
	stop=1
	go PrintMemUsage(&stop)
	time.Sleep(time.Microsecond)
	json_bench.JsonForUnmarshal(tblByte)
	atomic.StoreInt32(&stop,0)
}
func proto() {
	tblByte := json_bench.JsonFor()
	var stop int32
	stop=1
	go PrintMemUsage(&stop)
	time.Sleep(time.Microsecond)
	json_bench.JsonForUnmarshal(tblByte)
	atomic.StoreInt32(&stop,0)
}

func main()  {
	json()
	///runtime.GC()
	//time.Sleep(time.Second*5)
	//proto()
	return
}
