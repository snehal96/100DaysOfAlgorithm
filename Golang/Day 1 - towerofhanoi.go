package main

import (
	"fmt"
	"strconv"
)

func printHanoi(n int64, left, right, middle string) {
	if n > 0 {
		printHanoi(n-1, left, middle, right)
		fmt.Println(left + " => " + right)
		printHanoi(n-1, middle, right, left)
	}
}

func main() {
	var input string
	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println(err)
	}

	n, _ := strconv.ParseInt(input, 10, 64)
	printHanoi(n, "left", "right", "middle")
}
