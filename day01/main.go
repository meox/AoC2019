package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	masses := loadMass("puzzle.txt")
	var s float64
	for _, x := range masses {
		s += x
	}

	fmt.Printf("fuel: %d\n", fuel(s))
}

func fuel(mass float64) int64 {
	return int64(math.Floor(mass/3) - 2)
}

func loadMass(fileName string) []float64 {
	var rec []float64

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}
		rec = append(rec, f)
	}

	return rec
}
