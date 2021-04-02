package main

import (
	"encoding/json"
	"log"
)


type  Subdata struct {
	L1 map[string]interface{}
}

type Data struct{
	Ll Subdata 
	Nodes []string
}

type student struct {
	Name string
	Age int
	Ll Subdata
	Nodes []string
}

func main(){
	d:=map[string]interface{}{"aa":12}

	data := Data{Nodes:[]string{"a","b","c"},Ll:Subdata{L1:d}}
//s:=student{Name:"katechun",Age:35,Nodes:[]string{"a","b","c"},Ll:{"aa":12}}
	jsonBytes,err := json.Marshal(data)
	if err != nil {
		log.Fatal("err!@!!")
	}

	println(string(jsonBytes))
}
