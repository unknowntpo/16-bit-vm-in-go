package main

import (
	"fmt"
)

type memory []int

func newMemory(b int) memory {
	return make(memory, b)
}

func (m *memory) set(addr int, value int) {
	(*m)[addr] = value
}

// Get returns the value stored in memory address addr.
func (m *memory) get(addr int) int {
	return (*m)[addr]
}

func (m *memory) show() {
	fmt.Printf("[ ")
	for _, v := range *m {
		fmt.Printf("0x%x ", v)
	}
	fmt.Printf(" ]\n")
}
