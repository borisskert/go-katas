package kata

import (
	"strconv"
	"strings"
)

// SimpleAssembler https://www.codewars.com/kata/58e24788e24ddee28e000053/train/go
func SimpleAssembler(program []string) map[string]int {
	machine := NewMachine()
	machine.LoadProgram(program)
	machine.Execute()

	return machine.registers
}

type Machine struct {
	registers          map[string]int
	instructionPointer int
	instructions       []Instruction
}

func NewMachine() Machine {
	return Machine{
		make(map[string]int),
		0,
		[]Instruction{},
	}
}

func (m *Machine) LoadProgram(program []string) {
	instructions := make([]Instruction, 0)

	for _, line := range program {
		instruction := m.parseInstruction(line)
		instructions = append(instructions, instruction)
	}

	m.instructions = instructions
}

func (m *Machine) parseInstruction(line string) Instruction {
	tokens := strings.Fields(line)

	if tokens[0] == "mov" {
		return Mov{m, tokens[1], tokens[2]}
	}

	if tokens[0] == "inc" {
		return Inc{m, tokens[1]}
	}

	if tokens[0] == "dec" {
		return Dec{m, tokens[1]}
	}

	if tokens[0] == "jnz" {
		return Jnz{m, tokens[1], tokens[2]}
	}

	panic("Unknown instruction: " + line)
}

func (m *Machine) Execute() {
	for m.instructionPointer >= 0 && m.instructionPointer < len(m.instructions) {
		m.instructions[m.instructionPointer].execute()
		m.instructionPointer++
	}
}

type Instruction interface {
	execute()
}

type Mov struct {
	machine *Machine
	x       string
	y       string
}

func (m Mov) execute() {
	valueY, err := strconv.Atoi(m.y)

	if err != nil {
		valueY = m.machine.registers[m.y]
	}

	m.machine.registers[m.x] = valueY
}

type Inc struct {
	machine *Machine
	x       string
}

func (i Inc) execute() {
	i.machine.registers[i.x]++
}

type Dec struct {
	machine *Machine
	x       string
}

func (d Dec) execute() {
	d.machine.registers[d.x]--
}

type Jnz struct {
	machine *Machine
	x       string
	y       string
}

func (j Jnz) execute() {
	valueX, err := strconv.Atoi(j.x)

	if err != nil {
		valueX = j.machine.registers[j.x]
	}

	if valueX != 0 {
		valueY, err := strconv.Atoi(j.y)

		if err != nil {
			valueY = j.machine.registers[j.y]
		}

		j.machine.instructionPointer += valueY - 1
	}
}
