// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024

type Machine struct {
	P      uint64
	Memory []uint64
}

func New() *Machine {
	m := Machine{}
	m.Memory = make([]uint64, 1024)
	return &Machine{Memory: m.Memory}
}

// func (m *Machine) GetMemory(register uint64, memory []uint64) uint64 {

// 	return 0
// }
