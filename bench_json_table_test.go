package json_bench

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"runtime"
	"sync"
	"testing"
)
const tblNum = 1000

func JsonFor() [][]byte  {
	var tblByte [][]byte
	tbls := genTables(tblNum)
	for i :=0; i< tblNum; i++ {
		data, _ :=json.Marshal(tbls[i])
		tblByte = append(tblByte, data)
	}
	return tblByte
}
func JsoniterFor() [][]byte  {
	var tblByte [][]byte
	tbls := genTables(tblNum)
	for i :=0; i< tblNum; i++ {
		data, _ :=jsoniter.Marshal(tbls[i])
		tblByte = append(tblByte, data)
	}
	return tblByte
}

func JsonForUnmarshal(tblByte [][]byte)  {
	var tbl TableInfo
	var tblArray []*TableInfo
	for i :=0; i< tblNum; i++ {
		err := json.Unmarshal(tblByte[i], &tbl)
		if err != nil {
			fmt.Println("error")
			break
		}
		tblArray = append(tblArray, &tbl)
	}
	if len(tblArray) != tblNum {
		fmt.Println("error")
	}
}
func JsonForUnmarshalGC(tblByte [][]byte)  {
	var tbl TableInfo
	var tblArray []*TableInfo
	for i :=0; i< tblNum; i++ {
		err := json.Unmarshal(tblByte[i], &tbl)
		if err != nil {
			fmt.Println("error")
			break
		}
		tblArray = append(tblArray, &tbl)
		runtime.GC()
	}
	if len(tblArray) != tblNum {
		fmt.Println("error")
	}
}
func JsonForUnmarshalParallel(tblByte [][]byte)  {
	goNum := 10
	wg := &sync.WaitGroup{}
	wg.Add(goNum)
	for id := 0; id< goNum; id++ {
		wid :=id
		go func() {
			defer wg.Done()
			var tblArray []*TableInfo
			var tbl TableInfo
			for i :=wid; i< tblNum; i = i+goNum {
				err := json.Unmarshal(tblByte[i], &tbl)
				if err != nil {
					fmt.Println("error")
					break
				}
				tblArray = append(tblArray, &tbl)
			}
		}()
	}
	wg.Wait()
}
func JsonForUnmarshalParallelGC(tblByte [][]byte)  {
	goNum := 10
	wg := &sync.WaitGroup{}
	wg.Add(goNum)
	for id := 0; id< goNum; id++ {
		wid :=id
		go func() {
			defer wg.Done()
			var tblArray []*TableInfo
			var tbl TableInfo
			for i :=wid; i< tblNum; i = i+goNum {
				err := json.Unmarshal(tblByte[i], &tbl)
				if err != nil {
					fmt.Println("error")
					break
				}
				tblArray = append(tblArray, &tbl)
				runtime.GC()
			}
		}()
	}
	wg.Wait()
}
func JsoniterForUnmarshal(tblByte [][]byte)  {
	var tbl TableInfo
	for i :=0; i< tblNum; i++ {
		jsoniter.Unmarshal(tblByte[i], &tbl)
	}
}
func JsonForRangeUnmarshal(tblByte [][]byte)  {
	var tbl TableInfo
	for _, res := range tblByte {
		json.Unmarshal(res, &tbl)
	}
}
func JsoniterForRangeUnmarshal(tblByte [][]byte)  {
	var tbl TableInfo
	for _, res := range tblByte {
		jsoniter.Unmarshal(res, &tbl)
	}
}
func BenchmarkJsonForMarshal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		JsonFor()
	}
}
func BenchmarkJsoniterForMarshal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		JsoniterFor()
	}
}
func BenchmarkJsonForUnmarshalNoGC(b *testing.B) {
	tblByte := JsonFor()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		JsonForUnmarshal(tblByte)
	}
}
func BenchmarkJsonForUnmarshalParallelNoGC(b *testing.B) {
	tblByte := JsonFor()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		JsonForUnmarshalParallel(tblByte)
	}
}
func BenchmarkJsonForUnmarshalGC(b *testing.B) {
	tblByte := JsonFor()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		JsonForUnmarshalGC(tblByte)
	}
}
func BenchmarkJsonForUnmarshalParallelGC(b *testing.B) {
	tblByte := JsonFor()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		JsonForUnmarshalParallelGC(tblByte)
	}
}
func BenchmarkJsoniterForUnmarshal(b *testing.B) {
	tblByte := JsoniterFor()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		JsoniterForUnmarshal(tblByte)
	}
}
func BenchmarkJsonForRangeUnmarshal(b *testing.B) {
	tblByte := JsonFor()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		JsonForRangeUnmarshal(tblByte)
	}
}
func BenchmarkJsoniterForRangeUnmarshal(b *testing.B) {
	tblByte := JsoniterFor()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		JsoniterForRangeUnmarshal(tblByte)
	}
}