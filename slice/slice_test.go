package main

import (
	"math/rand"
	"testing"
	"time"
)






const size = 8192

//var d []string
//func init(){
//	d =genData(size)
//}


func genMap(size int)[]string{
	ret := make([]string,size)
	for i:=0;i<size;i++{
		ret[i]=randS(i+1)
	}
	return ret
}


func genSlice(size int)[]string{
	ret := make([]string,size)
	for i:=0;i<size;i++{
		ret=append(ret,randS(i+1))
	}
	return ret
}

const alphanum = "ABCDEFGHIJKLMNPOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
func randS(n int)string{
	r:=rand.New(rand.NewSource(time.Now().UnixNano()))
	var str string
	length := len(alphanum)
	for i:=0;i<n;i++{
		a:=alphanum[r.Intn(len(alphanum))%length]
		str +=string(a)
	}
	return str
}

func BenchmarkSlice(t *testing.B){

	for i:=0;i<t.N;i++{
		genSlice(size)
	}
}

func BenchmarkMap(t *testing.B){
	for i:=0;i<t.N;i++{
		genMap(size)
	}

}