package main

type memory []byte

func newMemory(b int) memory {
	return make(memory, b)
}

func (m *memory) Set(addr int, value byte) {
	(*m)[addr] = value
}

func (m *memory) Show() {
	fmt.Println(*m)
}
