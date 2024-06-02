package bus

type Bus interface {
	Read(addr uint16) uint8
	Write(addr uint16, value uint8)
	ReadWord(addr uint16) uint16
	WriteWord(addr uint16, value uint16)
}
