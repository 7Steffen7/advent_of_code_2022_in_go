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

func recursive_count(scanner *bufio.Scanner, dir string) (int, int) {
	var result, temp_counter, temp_result, counter int

	for scanner.Scan() {
		line := scanner.Text()
		split_str := strings.Split(line, " ")
		if split_str[0] == "$" && split_str[1] == "cd" && split_str[2] != ".." {
			temp_counter, temp_result = recursive_count(scanner, split_str[2])
			counter += temp_counter
			result += temp_result
		} else if split_str[0] == "$" && split_str[1] == "cd" && split_str[2] == ".." {
			fmt.Print("Dir:", dir)
			fmt.Print(" counter:", counter)
			if counter <= 100000 {
				result += counter
			}
			fmt.Println(" result:", result)
			return counter, result
		} else if isNumeric(split_str[0]) {
			temp, _ := strconv.Atoi(split_str[0])
			counter += temp
		}
	}
	fmt.Print("dir:", dir)
	fmt.Println(" counter:", counter)
	return counter, result
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	_ = scanner.Scan()
	_, result := recursive_count(scanner, "/")
	println("result", result)
}
