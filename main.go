package main

import "github.com/drewwalton19216801/butterfly-go/cpu6502"

func main() {
	bus := NewMainBus()
	cpu := cpu6502.NewCPU(bus)
	err := cpu.Reset()
	if err != nil {
		return
	}
	println(cpu.Status())
}
