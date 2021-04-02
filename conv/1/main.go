package main

import (
"fmt"
"math/big"
)

func toHexInt(n *big.Int) string {
	return fmt.Sprintf("%x", n) // or %X or upper case
}

func toHexRat(n *big.Rat) string {
	return fmt.Sprintf("%x", n) // or %X or upper case
}

func main() {
	a := big.NewInt(-59)
	b := big.NewInt(59)

	fmt.Printf("negative int lower case: %x\n", a)
	fmt.Printf("negative int upper case: %X\n", a) // %X or upper case

	fmt.Println("using Int function:", toHexInt(b))

	f := big.NewRat(3, 4) // fraction: 3/4

	fmt.Printf("rational lower case: %x\n", f)
	fmt.Printf("rational lower case: %X\n", f)

	fmt.Println("using Rat function:", toHexRat(f))
}
