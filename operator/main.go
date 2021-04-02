package main

import "fmt"

func main(){
	//fmt.Println(^1 )        // untyped integer constant, equal to -2
	////fmt.Println(uint8(^1))  // illegal: same as uint8(-2), -2 cannot be represented as a uint8
	//fmt.Println(^uint8(1) ) // typed uint8 constant, same as 0xFF ^ uint8(1) = uint8(0xFE)
	//fmt.Println(int8(^1)   )// same as int8(-2)
	//fmt.Println(^int8(1)  ) // same as -1 ^ int8(1) = -2

	for i:=0;i<10;i++{
	a := 1
	f := func() int { a++; return a }
	x := []int{a, f()}            // x may be [1, 2] or [2, 2]: evaluation order between a and f() is not specified
	//m := map[int]int{a: 1, a: 2}  // m may be {2: 1} or {2: 2}: evaluation order between the two map assignments is not specified
	//n := map[int]int{a: f()}      // n may be {2: 3} or {3: 3}: evaluation order between the key and the value is not specified


		fmt.Println(x)
		//fmt.Println(m)
		//fmt.Println(n)
	}

}
