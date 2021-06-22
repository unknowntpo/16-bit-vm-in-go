package main

import (
	"fmt"
)

type cpu struct {
	m memory
	r RegisterMap
}

func newCPU(mem memory) *cpu {
	return &cpu{m: mem, r: regMap}
}

func (c *cpu) step() error {
	ins, err := c.fetch()
	if err != nil {
		return fmt.Errorf("fail to step 1 operation: %v", err)
	}

	err = c.decodeAndExecute(Instruction(ins))
	if err != nil {
		return fmt.Errorf("fail to decode and execute instruction: %v", err)
	}

	return nil
}

// fetch fetches one instruction and return the value it fetched with error if occured.
func (c *cpu) fetch() (int, error) {
	nextInsAddr, err := c.r.getValue(REG_IP)
	if err != nil {
		return 0, fmt.Errorf("fail to get value from register %s: %v", REG_IP.String(), err)
	}

	value := c.m.get(nextInsAddr)
	c.r.setValue(REG_IP, nextInsAddr+1)
	return int(value), nil
}

// decodeAndExecute decodes the given instruction and execute with corresponding behavior.
func (c *cpu) decodeAndExecute(ins Instruction) error {
	switch ins {
	case MOV_LIT_R1:
		ub, err := c.fetch()
		if err != nil {
			return fmt.Errorf("MOV_LIT_R1 fail to fetch upper byte: %v", err)
		}
		lb, err := c.fetch()
		if err != nil {
			return fmt.Errorf("MOV_LIT_R1 fail to fetch lower byte: %v", err)
		}
		v := (ub << 8) | lb

		err = c.r.setValue(REG_R1, int(v))
		if err != nil {
			return fmt.Errorf("fail to set %s values: %v", REG_R1.String(), err)
		}
	case MOV_LIT_R2:
		ub, err := c.fetch()
		if err != nil {
			return fmt.Errorf("MOV_LIT_R2 fail to fetch upper byte: %v", err)
		}
		lb, err := c.fetch()
		if err != nil {
			return fmt.Errorf("MOV_LIT_R2 fail to fetch lower byte: %v", err)
		}
		v := (ub << 8) | lb

		err = c.r.setValue(REG_R2, int(v))
		if err != nil {
			return fmt.Errorf("fail to set %s values: %v", REG_R2.String(), err)
		}

	case ADD_REG_REG:
		reg1, err := c.fetch()
		if err != nil {
			return fmt.Errorf("ADD_REG_REG fail to fetch first register value: %v", err)
		}
		reg2, err := c.fetch()
		if err != nil {
			return fmt.Errorf("ADD_REG_REG fail to fetch second register value: %v", err)
		}

		valReg1, err := c.r.getValue(Register(reg1))
		if err != nil {
			return fmt.Errorf("ADD_REG_REG fail to get value from first register: %v", err)
		}

		valReg2, err := c.r.getValue(Register(reg2))
		if err != nil {
			return fmt.Errorf("ADD_REG_REG fail to get value from second register: %v", err)
		}

		err = c.r.setValue(REG_ACC, valReg1+valReg2)
		if err != nil {
			return fmt.Errorf("fail to set %s values: %v", REG_ACC.String(), err)
		}
	}
	return nil
}

func (c *cpu) show() {
	//c.m.show()
	c.r.show()
	fmt.Println()
}
