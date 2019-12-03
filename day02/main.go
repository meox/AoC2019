package main

import (
	"github.com/pkg/errors"
	"fmt"
)

const (
	ADD = 1
	MUL = 2
	HLT = 99
)

type Program []int

func IntCode(program Program) (Program, error) {
	addr := func(idx int) int {
		return program[idx]
	}

	value := func(idx int) int {
		return program[addr(idx)]
	}

	for idx := 0; idx < len(program); idx++ {
		op := program[idx]

		switch op {
		case ADD:
			a := value(idx + 1)
			b := value(idx + 2)
			program[addr(idx + 3)] = a + b
			idx += 3
		case MUL:
			a := value(idx + 1)
			b := value(idx + 2)
			program[addr(idx + 3)] = a * b
			idx += 3
		case HLT:
			return program, nil
		}
	}

	return program, errors.New("unexpected end")
}

func main() {
	prg1, err := IntCode([]int{2,4,4,5,99,0})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", prg1)

	src := []int{1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,9,19,1,19,5,23,2,6,23,27,1,6,27,31,2,31,9,35,1,35,6,39,1,10,39,43,2,9,43,47,1,5,47,51,2,51,6,55,1,5,55,59,2,13,59,63,1,63,5,67,2,67,13,71,1,71,9,75,1,75,6,79,2,79,6,83,1,83,5,87,2,87,9,91,2,9,91,95,1,5,95,99,2,99,13,103,1,103,5,107,1,2,107,111,1,111,5,0,99,2,14,0,0}
	// adjustment
	src[1] = 12
	src[2] = 2

	prg2, err := IntCode(src)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", prg2)
}
