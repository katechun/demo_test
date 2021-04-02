package main

import (
	"fmt"
)

func main(){
	a:="hello 世界!"
	fmt.Println(string(a[1]))
	fmt.Println(a[7])

	s:=[]int{1,2,3,4,5,6,7,8,9}
	ss:=s[1:3:5]
	ss=append(ss,[]int{11,22,33,44,55,66}...)
	ss[0]=111
	fmt.Println(s,ss)
}
