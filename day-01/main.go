package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var currNum, prevNum int
	increases := 0
	i := 0

	for scanner.Scan() {
		if i == 0 {
			prevNum, _ = strconv.Atoi(scanner.Text())
			i += 1
		}
		currNum, _ = strconv.Atoi(scanner.Text())
		if currNum > prevNum {
			increases += 1
		}
		prevNum = currNum
		i += 1
	}
	fmt.Printf("Number of increases: %d", increases)
}
