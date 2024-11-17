package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isNumeric(str string) bool {
	for _, value := range str {
		if !unicode.IsDigit(value) {
			return false
		}
	}
	return true
}

func recursive_count(scanner *bufio.Scanner, resultPTR *int, used_storage int) int {
	var counter, temp_count int

	for scanner.Scan() {
		line := scanner.Text()
		split_str := strings.Split(line, " ")
		if split_str[0] == "$" && split_str[1] == "cd" && split_str[2] != ".." {
			temp_count = recursive_count(scanner, resultPTR, used_storage)
			counter += temp_count
		} else if split_str[0] == "$" && split_str[1] == "cd" && split_str[2] == ".." {
			if 70000000-used_storage+counter >= 30000000 && counter < *resultPTR {
				*resultPTR = counter
			}
			return counter
		} else if isNumeric(split_str[0]) {
			temp, _ := strconv.Atoi(split_str[0])
			counter += temp
		}
	}
	return counter
}

func usedStorage() int {
	var disc_space int

	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split_line := strings.Split(line, " ")
		if isNumeric(split_line[0]) {
			temp, _ := strconv.Atoi(split_line[0])
			disc_space += temp
		}
	}
	return disc_space
}

func main() {
	used_storage := usedStorage()
	value := 70000000

	fmt.Println("used storage: ", used_storage)

	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	_ = scanner.Scan()
	recursive_count(scanner, &value, used_storage)
	println("result", value)

}
