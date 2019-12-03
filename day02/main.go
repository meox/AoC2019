package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strconv"
	"strings"
)

const (
	ADD = 1
	MUL = 2
	HLT = 99
)


type Program struct {
	ip int
	code []int

	// 1K RAM
	memory []int
}

func (p *Program) Load(fname string) error {
	file, err := os.Open(fname)
	if err != nil {
		return err
	}

	buff := make([]byte, 4096)
	n, err := file.Read(buff)
	if err != nil {
		return err
	}

	codes := strings.Split(string(buff[0:n]), ",")
	for _, e := range codes {
		opcode, err := strconv.Atoi(e)
		if err != nil {
			return err
		}
		p.code = append(p.code, opcode)
	}

	return nil
}

func (p *Program) Reset() {
	p.memory = make([]int, len(p.code))
	for idx, e := range p.code {
		p.memory[idx] = e
	}
}

func (p *Program) IntCode() error {
	addr := func(idx int) int {
		return p.memory[idx]
	}

	value := func(idx int) int {
		return p.memory[addr(idx)]
	}

	p.ip = 0
	for ; p.ip < len(p.memory); p.ip++ {
		op := p.memory[p.ip]

		switch op {
		case ADD:
			a := value(p.ip + 1)
			b := value(p.ip + 2)
			p.memory[addr(p.ip + 3)] = a + b
			p.ip += 3
		case MUL:
			a := value(p.ip + 1)
			b := value(p.ip + 2)
			p.memory[addr(p.ip + 3)] = a * b
			p.ip += 3
		case HLT:
			return nil
		}
	}

	return errors.New("unexpected end")
}

func main() {
	prg0 := Program{code: []int{2,4,4,5,99,0}}
	prg0.Reset()
	err := prg0.IntCode()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", prg0.code)

	var prg1 Program
	if err := prg1.Load("puzzle1.txt"); err != nil {
		panic(err)
	}
	prg1.Reset()
	// adjustment
	prg1.memory[1] = 12
	prg1.memory[2] = 2

	err = prg1.IntCode()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", prg1.code)

	var prg2 Program
	if err := prg2.Load("puzzle2.txt"); err != nil {
		panic(err)
	}

	for noun := 0; noun < 99; noun++ {
		for verb := 0; verb < 99; verb++ {
			prg2.Reset()

			// adjustment
			prg2.memory[1] = noun
			prg2.memory[2] = verb

			err = prg2.IntCode()
			if err != nil {
				panic(err)
			}

			if prg2.memory[0] == 19690720 {
				fmt.Printf("noun=%d, verb=%d, result=%d\n", noun, verb, 100 * noun + verb)
				goto End
			}
		}
	}
	End:
}
