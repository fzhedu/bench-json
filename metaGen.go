package json_bench

import (
	"encoding/json"
	"fmt"
	proto "github.com/gogo/protobuf/proto"
	jsoniter "github.com/json-iterator/go"
	"runtime"
	"sync"
)
const tblNum = 10000

func ProtoFor() [][]byte  {
	var tblByte [][]byte
	tbls := genPTables(tblNum)
	for i :=0; i< tblNum; i++ {
		data, _ :=proto.Marshal(tbls[i])
		tblByte = append(tblByte, data)
	}
	return tblByte
}
func ProtoForUnmarshal(tblByte [][]byte)  {
	var tblArray []*PTableInfo
	for i :=0; i< tblNum; i++ {
		var tbl PTableInfo
		err := proto.Unmarshal(tblByte[i], &tbl)
		if err != nil {
			fmt.Println("error")
			break
		}
		tblArray = append(tblArray, &tbl)
		if i % 100 == 0 {
			runtime.GC()
		}
	}
	if len(tblArray) != tblNum {
		fmt.Println("error")
	}
}


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
	var tblArray []*TableInfo
	for i :=0; i< tblNum; i++ {
		var tbl TableInfo
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
	var tblArray []*TableInfo
	for i :=0; i< tblNum; i++ {
		var tbl TableInfo
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
			for i :=wid; i< tblNum; i = i+goNum {
				var tbl TableInfo
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
func ProtoForUnmarshalParallel(tblByte [][]byte)  {
	goNum := 10
	wg := &sync.WaitGroup{}
	wg.Add(goNum)
	for id := 0; id< goNum; id++ {
		wid :=id
		go func() {
			defer wg.Done()
			var tblArray []*PTableInfo
			for i :=wid; i< tblNum; i = i+goNum {
				var tbl PTableInfo
				err := proto.Unmarshal(tblByte[i], &tbl)
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
			for i :=wid; i< tblNum; i = i+goNum {
				var tbl TableInfo
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