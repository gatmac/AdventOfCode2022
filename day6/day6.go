package main

import (
	"log"
	"os"
)

const inputFileName = "input.txt"
const markerLen = 14

// marker confirms the slice is all unique and return true. Else false.
func marker(slice []byte) bool {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] == slice[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	input, err := os.ReadFile(inputFileName)
	if err != nil {
		log.Panicf("Unable to read file %s: %v", inputFileName, err)
	}
	for i := 0; i < len(input)-markerLen; i++ {
		if marker(input[i : i+markerLen]) {
			log.Printf("Processed %d characters before reaching the first marker", i+markerLen)
			log.Println(string(input[i : i+markerLen]))
			os.Exit(0)
		}
	}
}
