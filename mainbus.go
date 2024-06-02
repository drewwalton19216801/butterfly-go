package main

type MainBus struct {
	memory []uint8
}

func NewMainBus() *MainBus {
	return &MainBus{
		memory: make([]uint8, 0x10000),
	}
}

func (bus *MainBus) Read(addr uint16) uint8 {
	return bus.memory[addr]
}

func (bus *MainBus) Write(addr uint16, value uint8) {
	bus.memory[addr] = value
}

func (bus *MainBus) ReadWord(addr uint16) uint16 {
	return uint16(bus.memory[addr]) | uint16(bus.memory[addr+1])<<8
}

func (bus *MainBus) WriteWord(addr uint16, value uint16) {
	bus.memory[addr] = uint8(value)
	bus.memory[addr+1] = uint8(value >> 8)
}

func (bus *MainBus) Reset() {
	for i := 0; i < len(bus.memory); i++ {
		bus.memory[i] = 0
	}
}
