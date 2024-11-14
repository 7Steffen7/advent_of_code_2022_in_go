package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func peekForward(subSlice []byte) bool {
	// fmt.Printf("%s \n", subSlice)
	for i := 0; i < len(subSlice); i++ {
		for j := 1; j+i < len(subSlice); j++ {
			if subSlice[i] == subSlice[i+j] {
				return false
			}
		}
	}
	return true
}

func main() {

	data, err := os.ReadFile("input.txt")
	check(err)

	// var input string = string(data)
	// fmt.Println(input)
	for i := 0; i+14 < len(data); i++ {
		if peekForward(data[i : i+14]) {
			fmt.Println("nbr:", i+14)
			break
		}
	}
}
