package memalign

import (
	"testing"
	"unsafe"
)

var ptrSize uintptr
func init(){
	ptrSize = unsafe.Sizeof(uintptr(1))
}
type SType struct {
	b [32]byte
}

func BenchmarkUnAligned(b *testing.B){
	x:=SType{}
	address := uintptr(unsafe.Pointer(&x.b))+1
	if address%ptrSize == 0 {
		b.Error("Not unaligned address")
	}
	tmp:=(*int64)(unsafe.Pointer(&x.b))
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		*tmp = int64(i)
	}
}


func BenchmarkAligned(b *testing.B){
	x:=SType{}
	address := uintptr(unsafe.Pointer(&x.b))
	if address%ptrSize != 0 {
		b.Error("Not aligned address")
	}
	tmp:=(*int64)(unsafe.Pointer(&x.b))
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		*tmp = int64(i)
	}
}


type Ag struct {
	sl []int16
	ptr *int64
	st struct {
		str string
	}
	m map[string]int16
	i interface{}
	arr [2]int8
	bl bool
}