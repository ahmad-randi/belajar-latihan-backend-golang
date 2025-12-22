package util

import "fmt"

func InputString(msg string) string {
	var input string
	fmt.Println(msg)
	fmt.Println(&input)
	return input
}

func InputInt(msg string) int {
	var input int
	fmt.Print(msg)
	fmt.Scanln(&input)
	return input
}
