package main

import (
	"fmt"
)

func main() {
	fmt.Printf("\t.globl main\n")
	fmt.Printf("main:\n")
	fmt.Printf("\tmov $42, %%rax\n")
	fmt.Printf("\tret\n")
}
