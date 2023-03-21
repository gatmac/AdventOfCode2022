// https://adventofcode.com/2022/day/3
package main

import (
	"bufio"
	"fmt"
	"os"
)

const inputFileName = "./input.txt"

func readInputFile() []string {
	file, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Error opening file!!!")
		panic(err)
	}
	defer file.Close()

	data := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		panic(err)
	}
	return data
}

func getDuplicate(bag string) rune {
	l := len(bag)
	for _, b1 := range bag[:l/2] {
		for _, b2 := range bag[l/2:] {
			if b1 == b2 {
				return b1
			}
		}
	}
	return '-'
}

func getPriority(data rune) int {
	var dVal int
	dRaw := int(data)
	if dRaw > 96 {
		dVal = (dRaw - 96)
	} else {
		dVal = (dRaw - (64 - 26))
	}
	return dVal
}

func getCommonItem(s1, s2, s3 string) rune {
	for _, r1 := range s1 {
		for _, r2 := range s2 {
			for _, r3 := range s3 {
				if r1 == r2 && r2 == r3 {
					return r1
				}
			}
		}
	}
	return '-'
}

func main() {
	var sum, commonSum int
	data := readInputFile()
	for n, d := range data {
		dup := getDuplicate((d))
		priority := getPriority(dup)
		fmt.Printf("%d %s: %c = %d\n", n, d, dup, priority)
		sum += priority

		// Part 2
		if n%3 == 0 {
			common := getCommonItem(data[n], data[n+1], data[n+2])
			fmt.Printf("The common letter of row %d is %c\n", n, common)
			commonSum += getPriority(common)
		}
	}
	fmt.Printf("a = %d, A = %d\n", int('a'), int('A'))
	fmt.Printf("The total priority is %d\n", sum)
	fmt.Printf("The priority of common items is %d\n", commonSum)
}
