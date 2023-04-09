package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFileName = "./input.txt"

/* func readInputFile() []string {
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
} */

func readInputCsv() [][]string {
	data := make([][]string, 0)

	// open file
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// do something with read line
		data = append(data, rec)
	}
	return data
}

type Range struct {
	Start int
	End   int
}

func (r Range) Contains(s Range) bool {
	if r.Start <= s.Start && r.End >= s.End {
		return true
	}
	return false
}

func (r Range) Overlaps(s Range) bool {
	if (s.Start >= r.Start && s.Start <= r.End) || (s.End >= r.Start && s.End <= r.End) {
		return true
	}
	return false
}

func parseRanges(input []string) []Range {
	rnge := make([]Range, 0)
	for _, i := range input {
		newRange := Range{}
		split := strings.Split(i, "-")
		a, err := strconv.Atoi(string(split[0]))
		if err != nil {
			log.Panicf("cannot be converted to int: %s", split[0])
		}
		newRange.Start = a
		if len(split) == 1 {
			newRange.End = a
		} else {
			b, err := strconv.Atoi(string(split[1]))
			if err != nil {
				log.Panicf("cannot be converted to int: %s", split[1])
			}
			newRange.End = b
		}
		rnge = append(rnge, newRange)
	}
	return rnge
}

func main() {
	var contains, overlaps int
	input := readInputCsv()
	for _, i := range input {
		r := parseRanges(i)
		//fmt.Printf("%s became %v\n", i, r)
		if r[0].Contains(r[1]) || r[1].Contains(r[0]) {
			contains += 1
		}
		if r[0].Overlaps(r[1]) || r[1].Overlaps(r[0]) {
			overlaps += 1
		} else {
			fmt.Printf("Does not overlap: %v\n", r)
		}
	}
	fmt.Printf("%d groups are contained\n", contains)
	fmt.Printf("%d groups are overlapping\n", overlaps)
}
