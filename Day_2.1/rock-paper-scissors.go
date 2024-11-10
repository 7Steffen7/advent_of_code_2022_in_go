package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		switch line {
		case "A X":
			sum += 3
		case "A Y":
			sum += 4
		case "A Z":
			sum += 8
		case "B X":
			sum += 1
		case "B Y":
			sum += 5
		case "B Z":
			sum += 9
		case "C X":
			sum += 2
		case "C Y":
			sum += 6
		case "C Z":
			sum += 7
		}
	}
	println("sum:", sum)
}
