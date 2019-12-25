package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetNumPass(start, stop int) (int, error) {
	regA, err := NewRegister(start)
	if err != nil {
		return 0, err
	}
	regB, err := NewRegister(stop)
	if err != nil {
		return 0, err
	}

	// find a starting point
	if !regA.IsValidPass() {
		regA = regA.Next()
	}

	nPass := 0
	for regA.IsLess(regB) {
		if regA.IsValidPass() {
			nPass++
		}
		regA = regA.Next()
	}

	return nPass, nil
}

func Parse(input string) (int, int, error) {
	tks := strings.Split(input, "-")
	a, err := strconv.Atoi(tks[0])
	if err != nil {
		return 0, 0, err
	}

	b, err := strconv.Atoi(tks[1])
	if err != nil {
		return 0, 0, err
	}

	return a, b, nil
}

func main() {
	start, stop, err := Parse("307237-769058")
	if err != nil {
		fmt.Println("Wrong input!")
		os.Exit(1)
	}

	np, err := GetNumPass(start, stop)
	if err != nil {
		panic(err)
	}
	fmt.Printf("num pass: %d", np)
}
