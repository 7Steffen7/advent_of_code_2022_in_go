package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func convertValue(value byte) int {
	if value >= 'a' && value <= 'z' {
		return int(value - 96)
	} else {
		return int(value - 38)
	}
}

func sharedLetter(first_str string, second_str string) byte {
	for i := 0; i < len(first_str); i++ {
		var temp_string string = string(first_str[i])
		if strings.Contains(second_str, temp_string) {
			return first_str[i]
		}
	}
	return 0
}

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
		string_size := len(line)

		first_string := line[:string_size/2]
		second_string := line[string_size/2:]
		letter := sharedLetter(first_string, second_string)
		temp := convertValue(letter)
		sum += temp
		fmt.Printf("letter: %+q", letter)
		fmt.Println("value: ", temp)
		// fmt.Println("shared letter:", letter)
		// fmt.Print("string: ", line)
		// fmt.Printf(" first: %s, second %s", first_string, second_string)
		// fmt.Println("string size:", string_size)
	}
	fmt.Println("sum:", sum)

}
