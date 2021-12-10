package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var currNum, prevNum, prevPrevNum int
	prevWindowSum := -1
	increases := 0
	i := 0

	for scanner.Scan() {
		if i == 0 {
			prevPrevNum, _ = strconv.Atoi(scanner.Text())
			i += 1
			continue
		}

		if i == 1 {
			prevNum, _ = strconv.Atoi(scanner.Text())
			i += 1
			continue
		}

		currNum, _ = strconv.Atoi(scanner.Text())
		windowSum := prevPrevNum + prevNum + currNum

		if windowSum > prevWindowSum && prevWindowSum != -1 {
			increases += 1
		}
		prevWindowSum = windowSum
		prevPrevNum = prevNum
		prevNum = currNum
		i += 1
	}
	fmt.Printf("Number of increases with 3 point window: %d\n", increases)
}
