package main

import (
	"fmt"
	"strconv"
	"strings"
)

//func main(){
//	x:="CD08"
//	d,_:=strconv.ParseInt(x, 16, 64)
//	fmt.Println(d)
//	fmt.Println(52488+13048)
//}


func hexaNumberToInteger(hexaString string) string {
	// replace 0x or 0X with empty String
	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	return numberStr
}

func main() {
	var hexaNumber string
	fmt.Print("Enter Hexadecimal Number:")
	fmt.Scanln(&hexaNumber)
	output, err := strconv.ParseInt(hexaNumberToInteger(hexaNumber), 16, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Output %d", output)
}