package json_bench

import (
	"encoding/json"
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

func JsonForUnmarshal(tblByte [][]byte)  {
	var tbl TableInfo
	for i :=0; i< tblNum; i++ {
		json.Unmarshal(tblByte[i], &tbl)
/*		if i==100 {
			fmt.Println(len(tblByte[i]))
		}*/
	}
}
func JsonForRangeUnmarshal(tblByte [][]byte)  {
	var tbl TableInfo
	for _, res := range tblByte {
		json.Unmarshal(res, &tbl)
		/*		if i==100 {
				fmt.Println(len(tblByte[i]))
			}*/
	}
}
func BenchmarkJsonForMarshal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		JsonFor()
	}
}
func BenchmarkJsonForUnmarshal(b *testing.B) {
	tblByte := JsonFor()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		JsonForUnmarshal(tblByte)
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