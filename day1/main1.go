// https://adventofcode.com/2022/day/1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const inputFileName = "./input.txt"
const N = 3

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

func main() {
	var max, subtotal, maxNTotal uint64
	var maxN [N]uint64

	input := readInputFile()

	for i := range input {
		if len(input[i]) > 0 {
			val, err := strconv.ParseUint(input[i], 10, 64)
			if err != nil {
				panic(err)
			}
			subtotal += val
		} else {
			if subtotal > max {
				max = subtotal
			}
			for j := 0; j < N; j++ {
				if subtotal > maxN[j] {
					for k := N - 1; k > j; k-- {
						maxN[k] = maxN[k-1]
					}
					maxN[j] = subtotal
					break
				}
			}
			fmt.Printf("Subtotal: %d, Max: %d, MaxN: %v\n", subtotal, max, maxN)
			subtotal = 0
		}
	}
	fmt.Printf("Max is %d\n", max)
	for i := 0; i < N; i++ {
		maxNTotal += maxN[i]
	}
	fmt.Printf("Total of top %d is %d\n", N, maxNTotal)
}
