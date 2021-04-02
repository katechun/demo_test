package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main(){
	goExecutable,_ := exec.LookPath("go")

	cmdGoVer := &exec.Cmd{
		Path: goExecutable,
		Args: []string{goExecutable,"version"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	if err := cmdGoVer.Run();err != nil {
		fmt.Println("Error:",err);
	}
}
