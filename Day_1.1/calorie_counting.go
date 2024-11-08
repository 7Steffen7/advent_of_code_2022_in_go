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

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var cur_elf_calories int
	var all_elves_calories []int

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Compare(line, "") == 0 {
			all_elves_calories = append(all_elves_calories, cur_elf_calories)
			cur_elf_calories = 0
			continue
		}
		line_int, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		cur_elf_calories += line_int
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	all_elves_calories = append(all_elves_calories, cur_elf_calories)
	slices.Sort(all_elves_calories)
	slices.Reverse(all_elves_calories)
	top_3 := all_elves_calories[:3]

	fmt.Println("top 3 calories", top_3[0]+top_3[1]+top_3[2])
}
