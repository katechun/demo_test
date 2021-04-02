package main

import (
"fmt"
"math/rand"
"time"
)

func generator(name string) <-chan string {
	c := make(chan string)
	go func() {
		v := 0
		for {
			time.Sleep(time.Duration(rand.Intn(3000)) * time.Microsecond)
			c <- fmt.Sprintf("%s sends message %d", name, v)
			v++
		}
	}()
	return c
}

//func fanIn(m ...<-chan string) <-chan string{
//	c := make(chan string)
//	for _, mm := range m {
//		go func(cc <-chan string) {
//			for {
//				c <- <-cc
//			}
//		}(mm)
//	}
//
//	return c
//}



func fanInBySelect1(m1, m2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case m := <-m1:
				c <- m
			case m := <-m2:
				c <- m
			}
		}
	}()
	return c
}

func fanInBySelect2(m1, m2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case c <- <- m1:
			case c <- <- m2:
			}
		}
	}()
	return c
}

func main() {
	m1 := generator("service1")
	m2 := generator("service2")
	m := fanInBySelect1(m1, m2)
	for mm := range m {
		fmt.Println(mm)
	}
}
