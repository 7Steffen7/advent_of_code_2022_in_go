package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func rowCount() int {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		i++
	}
	return i
}

func leftRight(scanner *bufio.Scanner, twoD [][]int8, visible_trees map[string]bool) int {
	total_count := 0

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		var current_nbr int8 = -1
		// current_nbr := 0
		for j, str_value := range line {
			value := int8(str_value) - 48
			if value > current_nbr {
				s := fmt.Sprint(i, "-", j)
				visible_trees[s] = true
				total_count++
				current_nbr = value
			}
			twoD[i] = append(twoD[i], value)
		}
	}
	return total_count
}

func rightLeft(twoD [][]int8, row_count int, visible_trees map[string]bool) int {
	count := 0

	for j := 0; j < row_count; j++ {
		var current_nbr int8 = -1
		for i := row_count - 1; i >= 0; i-- {
			if twoD[j][i] > current_nbr {
				s := fmt.Sprint(j, "-", i)
				visible_trees[s] = true
				current_nbr = twoD[j][i]
				count++
			}
		}
	}
	return count
}

func topButtom(twoD [][]int8, row_count int, visible_trees map[string]bool) int {
	count := 0

	for j := 0; j < row_count; j++ {
		var current_nbr int8 = -1
		for i := 0; i < row_count; i++ {
			if twoD[i][j] > current_nbr {
				s := fmt.Sprint(i, "-", j)
				visible_trees[s] = true
				current_nbr = twoD[i][j]
				count++
			}
		}
	}
	return count
}

func ButtomTop(twoD [][]int8, row_count int, visible_trees map[string]bool) int {
	count := 0

	for j := 0; j < row_count; j++ {
		var current_nbr int8 = -1
		for i := row_count - 1; i >= 0; i-- {
			if twoD[i][j] > current_nbr {
				s := fmt.Sprint(i, "-", j)
				visible_trees[s] = true
				current_nbr = twoD[i][j]
				count++
			}
		}
	}
	return count
}

func main() {
	row_count := rowCount()
	// total_count := 0
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	twoD := make([][]int8, row_count)
	visible_trees := make(map[string]bool)
	total_count := leftRight(scanner, twoD, visible_trees)
	total_count += rightLeft(twoD, row_count, visible_trees)
	total_count += topButtom(twoD, row_count, visible_trees)
	total_count += ButtomTop(twoD, row_count, visible_trees)

	// for i := 0; scanner.Scan(); i++ {
	// 	line := scanner.Text()
	// 	var current_nbr int8 = -1
	// 	// current_nbr := 0
	// 	for _, str_value := range line {
	// 		value := int8(str_value) - 48
	// 		if value > current_nbr {
	// 			total_count++
	// 			current_nbr = value
	// 		}
	// 		twoD[i] = append(twoD[i], value)
	// 	}
	// }
	// fmt.Println("twoD:", twoD)
	fmt.Println("total count:", total_count)
	// fmt.Println("visible trees", visible_trees)
	fmt.Println("visible trees:", len(visible_trees))

}
