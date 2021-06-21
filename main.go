package main

import (
	"fmt"
)

func main() {
	// Create memory with 256 bytes.
	memory := newMemory(256)
	// Init CPU with memory.
	cpu := newCPU(memory)

	// Set data and instructions in memory.
	memory.Set(0x0, MOV_LIT_R1)
	memory.Set(0x1, 0x12)
	memory.Set(0x2, 0x34)
	memory.Set(0x3, MOV_LIT_R2)
	memory.Set(0x4, 0xAB)
	memory.Set(0x5, 0xCD)
	memory.Set(0x6, ADD_REG_REG)
	memory.Set(0x7, 0x2) // r1
	memory.Set(0x8, 0x3) // r2

	cpu.Mem.Show()
	// Fetch a byte.
	//cpu.Step()

}
