package main

import (
	"fmt"
	"strconv"
)

func main() {
	doThis("Hello World!")
	rad := add(69, 96)
	dar := strconv.Itoa(rad)
	fmt.Println("The sum of 69 and 96 is : " + dar)
}

func doThis(input string) {
	fmt.Println(input)
}

func add(input1 int, input2 int) int {
	var darth = input1 + input2
	return darth
}
