package main

import (
	"fmt"
	"strings"
)

type Case struct {
	input string
}

func main() {

	cases := []Case{
		{
			input: "doni",
		},
		{
			input: "DONIINOD",
		},
		{
			input: "eFishery",
		},
		{
			input: "12345677654321",
		},
		{
			input: "Hello, world",
		},
		{
			input: "Goodbye",
		},
		{
			input: "Good Morning",
		},
		{
			input: "Branches",
		},
		{
			input: "Eats",
		},
	}

	for _, v := range cases {
		fmt.Println(isPalindromeUsingRune(v.input))
		fmt.Println(isPalindromeUsingArrays(v.input))
	}
}

func isPalindromeUsingRune(input string) bool {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func isPalindromeUsingArrays(input string) bool {
	arrStr := strings.Split(input, "")

	for i, j := 0, len(arrStr)-1; i < j; i, j = i+1, j-1 {
		fmt.Println("check 1: ", arrStr[i], "check 2: ", arrStr[j])
		if arrStr[i] != arrStr[j] {
			return false
		}
	}

	return true
}
