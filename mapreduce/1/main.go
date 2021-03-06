package main

import (
	"fmt"
	"strings"
)

func MapStrToStr(arr []string, fn func(s string) string) []string {
	var newArray = []string{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

func MapStrToInt(arr []string, fn func(s string) int) []int {
	var newArray = []int{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}

	return newArray
}

func Reduce(arr []string, fn func(s string) int) int {
	sum := 0
	for _, it := range arr {
		sum += fn(it)
	}
	return sum

}

func Filter(arr []int,fn func(n int)bool)[]int{
	var newArray = []int{}
	for _,it := range arr{
		if fn(it){
			newArray=append(newArray,it)
		}
	}
	return newArray
}

func main() {
	var list = []string{"Hao", "Chen", "MegaEase"}

	//map
	x := MapStrToStr(list, func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Printf("%v\n", x)

	//map
	y := MapStrToInt(list, func(s string) int {
		return len(s)
	})

	fmt.Println(y)

	//reduce
	z := Reduce(list,func(s string)int{
		return len(s)
	})
	fmt.Println(z)

	//filter
	var intset = []int{1,2,3,4,5,6,7,8,9,10}
	out := Filter(intset,func(n int)bool{
		return n%2 == 1
	})
	fmt.Println(out)

	out = Filter(intset,func(n int)bool{
		return n > 5
	})

	fmt.Println(out)

}
