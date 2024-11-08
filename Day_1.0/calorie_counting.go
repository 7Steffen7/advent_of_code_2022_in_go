package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	var max_elf_calories int

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Compare(line, "") == 0 {
			if cur_elf_calories > max_elf_calories {
				max_elf_calories = cur_elf_calories
			}
			// fmt.Println("current elf calories", cur_elf_calories)
			cur_elf_calories = 0
			continue
		}
		line_int, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		cur_elf_calories += line_int
		// fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	if cur_elf_calories > max_elf_calories {
		max_elf_calories = cur_elf_calories
	}
	// fmt.Println("current elf calories", cur_elf_calories)
	fmt.Println("max elf calories", max_elf_calories)
}
