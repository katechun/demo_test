package rpc

import "fmt"

func sayHi(user string)string{
	return fmt.Sprintf("Hi, '%s' :D",user)
}
