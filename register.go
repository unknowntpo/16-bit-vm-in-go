package main

import (
	"errors"
	"fmt"
)

type Register int

const (
	REG_IP = Register(iota) // 0x0
	REG_ACC
	REG_R1
	REG_R2
	REG_R3
	REG_R4
	REG_R5
	REG_R6
	REG_R7
	REG_R8
)

type RegisterMap map[Register]int

// regMap is where we store values to register.
var regMap = map[Register]int{
	REG_IP:  0,
	REG_ACC: 0,
	REG_R1:  0,
	REG_R2:  0,
	REG_R3:  0,
	REG_R4:  0,
	REG_R5:  0,
	REG_R6:  0,
	REG_R7:  0,
	REG_R8:  0,
}

var (
	ErrNoSuchRegister = errors.New("no such registers")
)

func (r Register) String() string {
	switch r {
	case REG_IP:
		return "ip"
	case REG_ACC:
		return "acc"
	case REG_R1:
		return "r1"
	case REG_R2:
		return "r2"
	case REG_R3:
		return "r3"
	case REG_R4:
		return "r4"
	case REG_R5:
		return "r5"
	case REG_R6:
		return "r6"
	case REG_R7:
		return "r7"
	case REG_R8:
		return "r8"
	default:
		return ""
	}
}

func (m RegisterMap) show() {
	for i := Register(0); i < Register(len(m)); i++ {
		fmt.Printf("%s: 0x%x\n", i.String(), m[i])
	}
}

// getValue returns value from given register index,
// if no such register, return -1 and error.
func (m RegisterMap) getValue(reg Register) (int, error) {
	if v, ok := m[reg]; !ok {
		return 0, ErrNoSuchRegister
	} else {
		return v, nil
	}
}

// setValue sets value to given register index,
// if no such register, return error.
func (m RegisterMap) setValue(reg Register, value int) error {
	if _, ok := m[reg]; ok {
		m[reg] = value
		return nil
	}
	return ErrNoSuchRegister
}
