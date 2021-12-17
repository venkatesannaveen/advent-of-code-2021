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
	var bitReadings [12]int
	var gammaRate, epsilonRate string

	for scanner.Scan() {
		currReading := scanner.Text()
		for i := 0; i < len(bitReadings); i++ {
			if currReading[i] == '0' {
				bitReadings[i] += 1
			} else if currReading[i] == '1' {
				bitReadings[i] -= 1
			}
		}
	}

	for i := 0; i < len(bitReadings); i++ {
		if bitReadings[i] > 0 {
			gammaRate += string('0')
			epsilonRate += string('1')
		} else {
			gammaRate += string('1')
			epsilonRate += string('0')
		}
	}

	gammaRateDec, _ := strconv.ParseInt(gammaRate, 2, 16)
	epsilonRateDec, _ := strconv.ParseInt(epsilonRate, 2, 16)

	fmt.Printf("Gamma Rate is: %v\n", gammaRateDec)
	fmt.Printf("Epsilon Rate is: %v\n", epsilonRateDec)
	fmt.Printf("Power Consumption is: %v\n", gammaRateDec*epsilonRateDec)

}
