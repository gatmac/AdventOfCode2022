// https://adventofcode.com/2022/day/2

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

func eval(opponent byte, me byte) (int, int) {
	var moveMap = map[byte]int{
		'A': 1,
		'B': 2,
		'C': 3,
		'X': 1,
		'Y': 2,
		'Z': 3,
	}
	oppScore := moveMap[opponent]
	myScore := moveMap[me]

	// Part 2
	myScore = (oppScore + myScore + 1) % 3
	if myScore == 0 {
		myScore = 3
	}
	fmt.Printf("%s: %d & %s: %d, ", string(opponent), oppScore, string(me), myScore)

	// Part 1
	if oppScore == myScore {
		return oppScore + 3, myScore + 3
	} else if oppScore == myScore-1 || oppScore == myScore+2 {
		return oppScore, myScore + 6
	} else {
		return oppScore + 6, myScore
	}

	/* if myScore == 1 {
		return oppScore + 6, myScore
	} else if myScore == 2 {
		return oppScore + 3, myScore + 3
	} else {
		return oppScore, myScore + 6
	} */
}

func main() {
	var oppScore, myScore int

	input := readInputFile()
	for row := range input {
		x, y := eval(input[row][0], input[row][2])
		oppScore += x
		myScore += y
	}

	fmt.Printf("Opposition score is %d\n", oppScore)
	fmt.Printf("My score is %d\n", myScore)
}
