package cpu6502

import (
	"github.com/drewwalton19216801/butterfly-go/cpu/bus"
	"strconv"
)

type CPU struct {
	bus bus.Bus

	a      uint8
	x      uint8
	y      uint8
	status uint8
	sp     uint8
	pc     uint16

	fetched uint8
	addrAbs uint16
	addrRel uint16
	// addrMode AddressingMode
	opcode uint8
	cycles uint8
	// currentInstruction Instruction
	currentInstructionString string
}

const (
	FlagNone = iota
	FlagCarry
	FlagZero
	FlagInterrupt
	FlagDecimal
	FlagBreak
	FlagUnused
	FlagOverflow
	FlagNegative
)

func NewCPU(bus bus.Bus) *CPU {
	return &CPU{
		bus: bus,
	}
}

func (cpu *CPU) setFlag(flag int, value bool) {
	if value {
		cpu.status |= 1 << flag
	} else {
		cpu.status &= ^(1 << flag)
	}
}

func (cpu *CPU) getFlag(flag int) bool {
	return (cpu.status & (1 << flag)) != 0
}

func (cpu *CPU) Tick() error {
	return nil
}

func (cpu *CPU) Reset() error {
	cpu.a = 0x00
	cpu.x = 0x00
	cpu.y = 0x00
	cpu.status = FlagNone | FlagUnused | FlagInterrupt
	cpu.sp = 0xFF
	cpu.pc = cpu.bus.ReadWord(0xFFFC)
	return nil
}

func (cpu *CPU) ToString() string {
	// Return a string representation of all registers (0x00-0xFF)
	return "A: 0x" + strconv.FormatUint(uint64(cpu.a), 16) +
		" X: 0x" + strconv.FormatUint(uint64(cpu.x), 16) +
		" Y: 0x" + strconv.FormatUint(uint64(cpu.y), 16) +
		" P: 0x" + strconv.FormatUint(uint64(cpu.status), 16) +
		" SP: 0x" + strconv.FormatUint(uint64(cpu.sp), 16) +
		" PC: 0x" + strconv.FormatUint(uint64(cpu.pc), 16)
}

func (cpu *CPU) Status() string {
	return cpu.ToString()
}

func (cpu *CPU) SetPC(addr uint16) {
	cpu.pc = addr
}

func (cpu *CPU) GetPC() uint16 {
	return cpu.pc
}

func (cpu *CPU) GetRegister(reg byte) uint8 {
	switch reg {
	case 'A':
		return cpu.a
	case 'X':
		return cpu.x
	case 'Y':
		return cpu.y
	default:
		return 0
	}
}

func (cpu *CPU) SetRegister(reg byte, value uint8) {
	switch reg {
	case 'A':
		cpu.a = value
	case 'X':
		cpu.x = value
	case 'Y':
		cpu.y = value
	}
}

func (cpu *CPU) GetStatus() uint8 {
	return cpu.status
}

func (cpu *CPU) SetStatus(status uint8) {
	cpu.status = status
}
