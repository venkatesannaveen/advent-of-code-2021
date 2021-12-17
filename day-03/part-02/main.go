package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func netZeroBit(nums []string, pos int) int {
	zeroBit := 0
	for _, val := range nums {
		if val[pos] == '0' {
			zeroBit += 1
		} else if val[pos] == '1' {
			zeroBit -= 1
		}
	}
	return zeroBit
}

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var numbers []string
	// var leastCommon, mostCommon string

	for scanner.Scan() {
		currReading := scanner.Text()
		numbers = append(numbers, currReading)
	}

	o2Rating := append([]string{}, numbers...)
	co2Rating := append([]string{}, numbers...)

	i := 0
	for {
		j := 0
		zeroBit := netZeroBit(o2Rating, i)
		var mostCommon rune

		if zeroBit > 0 {
			mostCommon = '0'
		} else {
			mostCommon = '1'
		}

		for {
			if o2Rating[j][i] != byte(mostCommon) {
				o2Rating = append(o2Rating[:j], o2Rating[j+1:]...)
				j -= 1
			}
			j += 1
			if j == len(o2Rating) {
				break
			}
		}
		if len(o2Rating) == 1 {
			break
		}
		i += 1
	}

	i = 0
	for {
		j := 0
		zeroBit := netZeroBit(co2Rating, i)
		var leastCommon rune

		if zeroBit > 0 {
			leastCommon = '1'
		} else {
			leastCommon = '0'
		}

		for {
			if co2Rating[j][i] != byte(leastCommon) {
				co2Rating = append(co2Rating[:j], co2Rating[j+1:]...)
				j -= 1
			}
			j += 1
			if j == len(co2Rating) {
				break
			}
		}
		if len(co2Rating) == 1 {
			break
		}
		i += 1
	}

	o2RatingDec, _ := strconv.ParseInt(o2Rating[0], 2, 16)
	co2RatingDec, _ := strconv.ParseInt(co2Rating[0], 2, 16)

	fmt.Printf("Oxygen Rating: %v\n", o2RatingDec)
	fmt.Printf("CO2 Rating: %v\n", co2RatingDec)
	fmt.Printf("Life Support Rating: %v\n", o2RatingDec*co2RatingDec)
}
