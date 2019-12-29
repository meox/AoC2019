package intcode

import (
	"errors"
	"fmt"
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

	input  int
	memory []int
}

func (p *Program) SetInput(input int) {
	p.input = input
}

func NewProgram(data string) (Program, error) {
	var p Program

	codes := strings.Split(data, ",")
	for _, e := range codes {
		inCode, err := strconv.Atoi(e)
		if err != nil {
			return p, err
		}
		p.code = append(p.code, inCode)
	}

	p.Reset()
	return p, nil
}

func (p *Program) Reset() {
	p.memory = make([]int, len(p.code))
	for idx, e := range p.code {
		p.memory[idx] = e
	}
}

func (p *Program) Run() error {
	p.ip = 0
	for p.ip < len(p.memory) {
		op := NewOpCode(p.memory[p.ip])

		switch op.Value {
		case ADD:
			a, err := op.ReadValueA(p.ip+1, p.memory)
			if err != nil {
				return fmt.Errorf("reading A parameter: %v", err)
			}

			b, err := op.ReadValueB(p.ip+2, p.memory)
			if err != nil {
				return fmt.Errorf("reading A parameter: %v", err)
			}

			err = op.StoreValueC(p.ip+3, p.memory, a+b)
			if err != nil {
				return err
			}
			p.ip += 4
		case MUL:
			a, err := op.ReadValueA(p.ip+1, p.memory)
			if err != nil {
				return fmt.Errorf("reading A parameter: %v", err)
			}

			b, err := op.ReadValueB(p.ip+2, p.memory)
			if err != nil {
				return fmt.Errorf("reading A parameter: %v", err)
			}

			err = op.StoreValueC(p.ip+3, p.memory, a*b)
			if err != nil {
				return err
			}
			p.ip += 4
		case STORE:
			err := op.StoreValueA(p.ip+1, p.memory, p.input)
			if err != nil {
				return err
			}
			p.ip += 2
		case PRINT:
			a, err := op.ReadValueA(p.ip+1, p.memory)
			if err != nil {
				return fmt.Errorf("reading A parameter: %v", err)
			}
			fmt.Println(a)
			p.ip += 2
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

func (op *OpCode) ModeParamA() int { return op.Mode[2] }
func (op *OpCode) ModeParamB() int { return op.Mode[1] }
func (op *OpCode) ModeParamC() int { return op.Mode[0] }

func (op *OpCode) ReadValueA(idx int, memory []int) (int, error) {
	return op.readValue(op.ModeParamA(), idx, memory)
}

func (op *OpCode) ReadValueB(idx int, memory []int) (int, error) {
	return op.readValue(op.ModeParamB(), idx, memory)
}

func (op *OpCode) ReadValueC(idx int, memory []int) (int, error) {
	return op.readValue(op.ModeParamC(), idx, memory)
}

func (op *OpCode) StoreValueA(idx int, memory []int, val int) error {
	return op.storeValue(op.ModeParamA(), idx, memory, val)
}

func (op *OpCode) StoreValueB(idx int, memory []int, val int) error {
	return op.storeValue(op.ModeParamB(), idx, memory, val)
}

func (op *OpCode) StoreValueC(idx int, memory []int, val int) error {
	return op.storeValue(op.ModeParamC(), idx, memory, val)
}

func (op *OpCode) readValue(mode int, idx int, memory []int) (int, error) {
	switch mode {
	case IMMEDIATE_MODE:
		return memory[idx], nil
	case POSITION_MODE:
		addr := memory[idx]
		return memory[addr], nil
	}
	return 0, errors.New("io error")
}

func (op *OpCode) storeValue(mode int, idx int, memory []int, val int) error {
	switch mode {
	case IMMEDIATE_MODE:
		memory[idx] = val
	case POSITION_MODE:
		addr := memory[idx]
		memory[addr] = val
	default:
		return errors.New("mode not supported")
	}
	return nil
}
