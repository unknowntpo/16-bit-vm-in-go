package main

func main() {
	// Create memory with 256 bytes.
	memory := newMemory(256)
	// Init CPU with memory.
	cpu := newCPU(memory)

	// Set data and instructions in memory.
	memory.set(0x0, int(MOV_LIT_R1))
	memory.set(0x1, 0x12)
	memory.set(0x2, 0x34)
	memory.set(0x3, int(MOV_LIT_R2))
	memory.set(0x4, 0xAB)
	memory.set(0x5, 0xCD)
	memory.set(0x6, int(ADD_REG_REG))
	memory.set(0x7, int(REG_R1)) // r1
	memory.set(0x8, int(REG_R2)) // r2

	cpu.show()
	// Fetch a byte.
	// MOV_LIT_R1 0x1234 REG_R1
	cpu.step()
	cpu.show()

	// MOV_LIT_R2 0xABCD REG_R2
	cpu.step()
	cpu.show()
	// ADD_REG_REG REG_R1 REG_R2
	cpu.step()
	cpu.show()

}
