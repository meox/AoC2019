package main

import (
	"io/ioutil"
	"os"

	"github.com/meox/aoc2019/day5/intcode"
)

func main() {
	file, err := os.Open("puzzle.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buff, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	p, err := intcode.NewProgram(string(buff))
	if err != nil {
		panic(err)
	}

	p.SetInput(5)
	err = p.Run()
	if err != nil {
		panic(err)
	}
}
