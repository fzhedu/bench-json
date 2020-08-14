package json_bench

import "testing"

//func BenchmarkJsonForMarshal(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		JsonFor()
//	}
//}
//func BenchmarkJsoniterForMarshal(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		JsoniterFor()
//	}
//}
//func BenchmarkProtoForMarshal(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < 1; i++ {
//		ProtoFor()
//	}
//}
func BenchmarkJsonForUnmarshalNoGC(b *testing.B) {
	tblByte := JsonFor()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		JsonForUnmarshal(tblByte)
	}
}
func BenchmarkProtoForUnmarshalNoGC(b *testing.B) {
	tblByte := ProtoFor()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ProtoForUnmarshal(tblByte)
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
func BenchmarkProteForUnmarshalParallelNoGC(b *testing.B) {
	tblByte := ProtoFor()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ProtoForUnmarshalParallel(tblByte)
	}
}
//func BenchmarkJsonForUnmarshalGC(b *testing.B) {
//	tblByte := JsonFor()
//	b.ResetTimer()
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		JsonForUnmarshalGC(tblByte)
//	}
//}
//func BenchmarkJsonForUnmarshalParallelGC(b *testing.B) {
//	tblByte := JsonFor()
//	b.ResetTimer()
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		JsonForUnmarshalParallelGC(tblByte)
//	}
//}
//func BenchmarkJsoniterForUnmarshal(b *testing.B) {
//	tblByte := JsoniterFor()
//	b.ResetTimer()
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		JsoniterForUnmarshal(tblByte)
//	}
//}
//func BenchmarkJsonForRangeUnmarshal(b *testing.B) {
//	tblByte := JsonFor()
//	b.ResetTimer()
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		JsonForRangeUnmarshal(tblByte)
//	}
//}
//func BenchmarkJsoniterForRangeUnmarshal(b *testing.B) {
//	tblByte := JsoniterFor()
//	b.ResetTimer()
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		JsoniterForRangeUnmarshal(tblByte)
//	}
//}