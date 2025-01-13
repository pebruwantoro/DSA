package main

import "fmt"

type Case struct {
	input int
}

func main() {

	cases := []Case{
		{
			input: 121,
		},
		{
			input: 1441,
		},
		{
			input: 123456789009876,
		},
		{
			input: 12345677654321,
		},
		{
			input: 1418401,
		},
		{
			input: 904810413,
		},
		{
			input: 988917,
		},
		{
			input: 1200,
		},
		{
			input: 10000,
		},
	}

	for _, v := range cases {
		fmt.Println(isPalindrome(v.input))
	}
}

func isPalindrome(input int) bool {
	reversed := 0
	check := input

	if check < 0 {
		return false
	}

	for check != 0 {
		modulo := check % 10
		check /= 10
		reversed = reversed*10 + modulo
	}

	fmt.Println("reversed: ", reversed, "input: ", input)
	return input == reversed
}
