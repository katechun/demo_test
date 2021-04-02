package main

import "fmt"

func main() {
	slice := []int{0, 11, 22, 33}
	myMap := make(map[int]*int)

	for index, value := range slice {
		myMap[index] = &value
	}
	fmt.Println("=====new map=====")
	for k, v := range myMap {
		fmt.Printf("%d => %d\n", k, *v)
	}


	v := []int{1, 2, 3}
	for i := range v {
		fmt.Println(i)
		v = append(v, i)
	}
}
