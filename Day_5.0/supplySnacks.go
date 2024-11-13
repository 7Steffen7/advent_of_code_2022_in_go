package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func isAlpha(char byte) bool {
	if char >= 'A' && char <= 'Z' {
		return true
	}
	return false
}

func createStacks(arr []string, count int) [][]byte {
	stack := make([][]byte, 9)
	count--
	for i := 0; count-i > -1; i++ {
		row := arr[count-i]
		var idx int
		for i := 1; i < len(row); i += 4 {
			if isAlpha(row[i]) {
				stack[idx] = append(stack[idx], row[i])
			}
			idx++
			// fmt.Printf("% q", row[i])
		}
		// fmt.Print("\n")
	}
	return stack
}

func reorderStacks(scanner *bufio.Scanner, stack [][]byte) {
	line_split := make([]string, 6)
	var nbrs_to_move, source, dest int
	// for idx, value := range stack {
	// 	fmt.Printf("%d % s \n", idx+1, value)
	// }
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		line_split = strings.Split(line, " ")
		nbrs_to_move, _ = strconv.Atoi(line_split[1])
		source, _ = strconv.Atoi(line_split[3])
		dest, _ = strconv.Atoi(line_split[5])
		if nbrs_to_move > len(stack[source-1]) {
			nbrs_to_move = len(stack[source-1])
		}
		temp := stack[source-1][len(stack[source-1])-nbrs_to_move:]
		// fmt.Printf("[%s]\n", temp)
		slices.Reverse(temp)
		stack[dest-1] = append(stack[dest-1], temp...)
		stack[source-1] = stack[source-1][:len(stack[source-1])-nbrs_to_move]

		// for idx, value := range stack {
		// 	fmt.Printf("%d % s \n", idx+1, value)
		// }
	}
	for _, value := range stack {
		if len(value) != 0 {
			fmt.Printf("%c", value[len(value)-1])
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var arr []string
	var count int

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.ContainsAny(line, "[") {
			break
		}
		count++
		// fmt.Println(line)
		arr = append(arr, line)
	}
	stack := createStacks(arr, count)

	scanner.Scan()
	reorderStacks(scanner, stack)
	fmt.Print("\n")
}
