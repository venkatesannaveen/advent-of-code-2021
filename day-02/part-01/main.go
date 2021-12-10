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
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var horizontal, depth int

	for scanner.Scan() {
		currMove := strings.Fields(scanner.Text())
		steps, _ := strconv.Atoi(currMove[1])

		if currMove[0] == "forward" {
			horizontal += steps
		}

		if currMove[0] == "up" {
			depth -= steps
		}

		if currMove[0] == "down" {
			depth += steps
		}
	}

	fmt.Printf("Horizontal Steps: %v\n", horizontal)
	fmt.Printf("Depth Steps: %v\n", depth)
	fmt.Printf("Multiplied horizontal and depth: %v\n", horizontal*depth)
}
