package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	ADD   = 1
	MUL   = 2
	STORE = 3
	PRINT = 4
	HLT   = 99
)

type Program struct {
	ip   int
	code []int

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
			p.memory[addr(p.ip+3)] = a + b
			p.ip += 3
		case MUL:
			a := value(p.ip + 1)
			b := value(p.ip + 2)
			p.memory[addr(p.ip+3)] = a * b
			p.ip += 3
		case HLT:
			return nil
		}
	}

	return errors.New("unexpected end")
}
