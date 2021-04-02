package main

import (
	"fmt"
	"crypto/sha256"
	"encoding/hex"
)

func main(){
	var nodes = map[string]string{}
	node := "192.168.1.113"
	genHash := GenSha256(node)
	nodes[genHash] = node
	fmt.Println(nodes)

}


func GenSha256(s string)string {
	sum:= sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}
