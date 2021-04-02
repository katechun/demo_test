package main


func main(){
	defer println("defer 1")
	level0()
	println("main")
}


func level0(){
	level1()
	println("level0")
}

func level1(){
	defer println("defer func 3")
	defer func(){
		if err := recover();err != nil {
			println("recovering in progress...")
		}
	}()

	defer println("defer 2")

	level2()
	println("level1")
}

func level2(){
	defer println("defer func 4")
	panic("error level2")
	println("level2")
}
