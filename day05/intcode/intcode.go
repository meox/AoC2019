package intcode

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

const (
	ADD   = 1
	MUL   = 2
	STORE = 3
	PRINT = 4
	HLT   = 99

	POSITION_MODE  = 0
	IMMEDIATE_MODE = 1
)

type OpCode struct {
	Value int
	Mode  [3]int
}

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
		inCode, err := strconv.Atoi(e)
		if err != nil {
			return err
		}
		p.code = append(p.code, inCode)
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

func NewOpCode(code int) OpCode {
	switch code {
	case ADD, MUL, STORE, PRINT, HLT:
		return OpCode{
			Value: code,
		}
	default:
		var inCode [2]int
		var opCode OpCode
		j := 5
		for code > 0 {
			v := code % 10
			if j > 3 {
				inCode[j-4] = v
			} else {
				opCode.Mode[j-1] = v
			}

			code /= 10
			j--
		}
		opCode.Value = inCode[0]*10 + inCode[1]
		return opCode
	}
}

func (op *OpCode) A() int { return op.Mode[0] }
func (op *OpCode) B() int { return op.Mode[1] }
func (op *OpCode) C() int { return op.Mode[2] }
