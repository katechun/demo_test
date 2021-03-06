package main
import "fmt"

func main() {

forLoop:for number := 1; number < 10; number++ {
	fmt.Printf("%d", number)
	switch {
	case number == 1:
		fmt.Println("-- One")
	case number == 2:
		fmt.Println("-- Two")
	case number == 3:
		fmt.Println("-- Three")
	case number == 4:
		fmt.Println("-- Four")
	case number == 5:
		fmt.Println("-- Five")
	case number == 6:
		fmt.Println("-- Six")
	case number > 2:
		fmt.Println("-- Greater than two")
		break forLoop
	case number == 8:
		fmt.Println("-- Eight")
	case number == 9:
		fmt.Println("-- Nine")
	default:
		fmt.Println("-- Number not identified")
	}
}

fmt.Println("The switch outside!")
}