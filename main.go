package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	var source []byte
	source , _ = ioutil.ReadFile("/dev/stdin")
	number , err := strconv.Atoi(string(source))
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("\t.globl main\n")
	fmt.Printf("main:\n")
	fmt.Printf("\tmov $%d, %%rax\n", number)
	fmt.Printf("\tret\n")
}
