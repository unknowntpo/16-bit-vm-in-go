package main

type Instruction int

const (
	MOV_LIT_R1 = Instruction(0x10 + iota)
	MOV_LIT_R2
	ADD_REG_REG
)
