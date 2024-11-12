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
	nbrsPair := make([]int, 4)
	var counter int

	for scanner.Scan() {
		line := scanner.Text()
		prep := strings.Replace(line, ",", "-", 1)
		split_strings := strings.Split(prep, "-")

		for i := 0; i < 4; i++ {
			nbrsPair[i], err = strconv.Atoi(split_strings[i])
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Printf("nbrPairs: %d %d", nbrsPair[:2], nbrsPair[2:])
		if nbrsPair[2] <= nbrsPair[1] && nbrsPair[3] >= nbrsPair[1] {
			fmt.Print(" overlaps")
			counter++
		} else if nbrsPair[0] <= nbrsPair[2] && nbrsPair[1] >= nbrsPair[2] {
			fmt.Print(" overlaps")
			counter++
		} else if nbrsPair[3] >= nbrsPair[0] && nbrsPair[3] <= nbrsPair[1] {
			fmt.Print(" overlaps")
			counter++
		}
		print("\n")
	}
	fmt.Println("Count:", counter)
}
