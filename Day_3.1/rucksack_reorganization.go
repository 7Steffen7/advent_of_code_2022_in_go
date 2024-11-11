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

func getSharedValue(arr []string) byte {
	for i := 0; i < len(arr[0]); i++ {
		temp_string := string(arr[0][i])
		if strings.Contains(arr[1], temp_string) && strings.Contains(arr[2], temp_string) {
			return arr[0][i]
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
	s := make([]string, 3)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		switch i {
		case 0:
			s[0] = line
		case 1:
			s[1] = line
		case 2:
			s[2] = line
			value := getSharedValue(s)
			sum += convertValue(value)
			i = -1
		}
	}
	fmt.Println("sum:", sum)

}
