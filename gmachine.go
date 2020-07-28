// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024

type Mem struct {
	P      uint64
	Memory []uint64
}

func New() *Mem {
	return &Mem{}
}

func (m *Mem) GetMemory(in uint64) uint64 {
	return 0
}
