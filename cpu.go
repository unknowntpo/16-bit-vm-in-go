package main

import "fmt"

type INSTRUCTION byte

const (
	MOV_LIT_R1 = INSTRUCTION(0x10 + iota)
	MOV_LIT_R2
	MOV_REG_REG
)

var reg = map[string]int{
	"ip":  0,
	"acc": 0,
	"r1":  0,
	"r2":  0,
	"r3":  0,
	"r4":  0,
	"r5":  0,
	"r6":  0,
	"r7":  0,
	"r8":  0,
}

type cpu struct {
	m memory
	r reg
}
