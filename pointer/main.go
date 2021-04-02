package main

import "fmt"

func main(){
	//a:=[]int{1,2,3,4,5,6}
	//fmt.Println(&a)
	//fmt.Println(&a[0])
	//fmt.Println(&a[1])
	//fmt.Println(&a[2])
	//fmt.Println(&a[0] == &a[0])


	//fmt.Println(string([]byte{'a'}))      // not a constant: []byte{'a'} is not a constant
	//fmt.Println((*int)(nil)   )           // not a constant: nil is not a constant, *int is not a boolean, numeric, or string type

	fmt.Println(string('a') )      // "a"
	fmt.Println(string(-1) )       // "\ufffd" == "\xef\xbf\xbd"
	fmt.Println(string(0xf8) )     // "\u00f8" == "ø" == "\xc3\xb8"
	type MyString string
	fmt.Println(MyString(0x65e5) ) // "\u65e5" == "日" == "\xe6\x97\xa5"
}
