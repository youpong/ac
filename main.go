package main

import (
	"fmt"
	"io/ioutil"
	"errors"
)

type Token struct {
	kind string // "intliteral", "punct"
	value string
}

var source []byte
var sourceIndex = 0

func getChar() (byte, error) {
	if sourceIndex == len(source) {
		return 0, errors.New("EOF")
	}
	char := source[sourceIndex]
	sourceIndex++
	return char, nil
}

func ungetChar() {
	sourceIndex--
}

func readNumber(char byte) string {
	var number []byte = []byte{char}
	for {
		char, err := getChar()
		if err != nil {
			break
		}
		if '0' <= char && char <= '9' {
			number = append(number, char)
		} else {
			ungetChar()
			break
		}
	}
	return string(number)
}

func tokenize() []*Token {
	fmt.Printf("# tokens = ")
	var tokens []*Token
	for {
		char , err := getChar()
		if err != nil {
			break
		}
		switch char {
		case ' ', '\n':
			continue
		case '0','1','2','3','4','5','6','7','8','9':
			number := readNumber(char)
			token := &Token{
				kind: "intliteral",
				value: number,
			}
			tokens = append(tokens, token)
			fmt.Printf(" '%s' ", token.value)
		case '+', '-', ';':
			token := &Token{
				kind: "punct",
				value: string([]byte{char}),
			}
			tokens = append(tokens, token)
			fmt.Printf(" '%s' ", token.value)
		default:
			
			panic(fmt.Sprintf("Invalid char '%c'", char))
		}
	}
	fmt.Printf("\n")
	return tokens
}

func main() {
	source , _ = ioutil.ReadFile("/dev/stdin")
	tokenize()
}
