package cpu6502

type AddressingMode int

const (
	Accumulator AddressingMode = iota
	Immediate
	Implied
	Indirect
	Relative
	Absolute
	IndexedIndirect
	IndirectIndexed
	AbsoluteX
	AbsoluteY
	ZeroPage
	ZeroPageX
	ZeroPageY
)

var AddressingModes = map[AddressingMode]string{
	Accumulator:     "Accumulator",
	Immediate:       "Immediate",
	Implied:         "Implied",
	Indirect:        "Indirect",
	Relative:        "Relative",
	Absolute:        "Absolute",
	IndexedIndirect: "IndexedIndirect",
	IndirectIndexed: "IndirectIndexed",
	AbsoluteX:       "AbsoluteX",
	AbsoluteY:       "AbsoluteY",
	ZeroPage:        "ZeroPage",
	ZeroPageX:       "ZeroPageX",
	ZeroPageY:       "ZeroPageY",
}

func (mode AddressingMode) String() string {
	return AddressingModes[mode]
}

// executeAddressingMode executes the specified addressing mode and returns the number of extra cycles.
//
// Parameters:
// - mode: The addressing mode to execute.
//
// Returns:
// - int: The number of extra cycles required for the addressing mode.
func (cpu *CPU) executeAddressingMode(mode AddressingMode) int {
	switch mode {
	case Accumulator:
		return 0 // Nothing to do
	case Immediate:
		cpu.addrAbs = cpu.pc
		cpu.pc++
		return 0
	case Implied:
		cpu.fetched = cpu.a
		return 0
	case Indirect:
		addrLo := cpu.bus.Read(cpu.pc)
		addrHi := cpu.bus.Read(cpu.pc + 1)
		address := uint16(addrHi)<<8 | uint16(addrLo)
		if addrLo == 0x00FF {
			lowByte := cpu.bus.Read(address & 0xFF00)
			highByte := cpu.bus.Read(address)
			cpu.addrAbs = uint16(highByte)<<8 | uint16(lowByte)
		} else {
			lowByte := cpu.bus.Read(address)
			highByte := cpu.bus.Read(address + 1)
			cpu.addrAbs = uint16(highByte)<<8 | uint16(lowByte)
		}
		cpu.pc += 2
		return 0
	case Relative:
		cpu.addrRel = uint16(cpu.bus.Read(cpu.pc))
		cpu.pc++
		if cpu.addrRel&0x80 != 0 {
			cpu.addrRel |= 0xFF00
		}
		return 0
	case Absolute:
		address := cpu.bus.ReadWord(cpu.pc)
		cpu.addrAbs = address
		cpu.pc += 2
		return 0
	case IndexedIndirect:
		temp := cpu.bus.Read(cpu.pc)
		lowByte := cpu.bus.Read(uint16(temp + cpu.x))
		highByte := cpu.bus.Read(uint16(temp + cpu.x + 1))
		cpu.addrAbs = uint16(highByte)<<8 | uint16(lowByte)
		cpu.pc++
		return 0
	case IndirectIndexed:
		temp := cpu.bus.Read(cpu.pc)
		lowByte := cpu.bus.Read(uint16(temp))
		highByte := cpu.bus.Read(uint16(temp + 1))
		cpu.addrAbs = uint16(highByte)<<8 | uint16(lowByte)
		cpu.pc++
		// If we crossed a page boundary, we need to add an extra cycle
		if (cpu.addrAbs & 0xFF00) != (cpu.pc & 0xFF00) {
			return 1
		}
		return 0
	case AbsoluteX:
		address := cpu.bus.ReadWord(cpu.pc)
		cpu.addrAbs = address + uint16(cpu.x)
		cpu.pc += 2
		// If we crossed a page boundary, we need to add an extra cycle
		if (cpu.addrAbs & 0xFF00) != (address & 0xFF00) {
			return 1
		}
		return 0
	case AbsoluteY:
		address := cpu.bus.ReadWord(cpu.pc)
		cpu.addrAbs = address + uint16(cpu.y)
		cpu.pc += 2
		// If we crossed a page boundary, we need to add an extra cycle
		if (cpu.addrAbs & 0xFF00) != (address & 0xFF00) {
			return 1
		}
		return 0
	case ZeroPage:
		cpu.addrAbs = uint16(cpu.bus.Read(cpu.pc) & 0x00FF)
		cpu.pc++
		return 0
	case ZeroPageX:
		cpu.addrAbs = uint16(cpu.bus.Read(cpu.pc)&0x00FF) + uint16(cpu.x)
		cpu.pc++
		return 0
	case ZeroPageY:
		cpu.addrAbs = uint16(cpu.bus.Read(cpu.pc)&0x00FF) + uint16(cpu.y)
		cpu.pc++
		return 0
	default:
		return 0
	}
}

func (cpu *CPU) GetAbsoluteAddress() uint16 {
	return cpu.addrAbs
}

func (cpu *CPU) GetRelativeAddress() uint16 {
	return cpu.addrRel
}
