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
	var s int64
	for _, x := range masses {
		fuelForMass := fuel(x)
		correction := calcCorrectFuel(float64(fuelForMass))
		s += fuelForMass + correction
	}

	fmt.Printf("fuel: %d\n", s)
}

func calcCorrectFuel(mass float64) int64 {
	var f int64
	for {
		v := fuel(mass)
		if v <= 0 {
			break
		}
		f += v
		mass = float64(v)
	}
	return f
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
