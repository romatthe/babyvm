package main

import "fmt"

// All available instructions for our VM
const (
	PSH = iota
	ADD
	POP
	SET
	HLT
)

// All available registers for our VM
const (
	A = iota
	B
	C
	D
	E
	F
	NumOfRegisters
)

var program = [...]uint8{
	PSH, 5,
	PSH, 6,
	ADD,
	POP,
	HLT}

type VM struct {
	running   bool
	pc        uint32
	stack     stack
	registers [NumOfRegisters]uint8
}

func NewVM() *VM {
	return &VM{
		running:   true,
		pc:        0,
		stack:     make(stack, 0),
		registers: [NumOfRegisters]uint8{},
	}
}

func (vm *VM) fetch() uint8 {
	var inst = program[vm.pc]
	vm.pc++
	return inst
}

func (vm *VM) eval(inst uint8) {
	switch inst {
	case HLT:
		vm.running = false
		fmt.Println("Halting!")
	case PSH:
		vm.stack = vm.stack.Push(vm.fetch())
	case POP:
		var s, p = vm.stack.Pop()
		vm.stack = s
		fmt.Printf("Popped %v off the stack\n", p)
	case ADD:
		var s, p = vm.stack.Pop()
		var t, q = s.Pop()
		vm.stack = t.Push(p + q)
	default:
		fmt.Println("Unrecognized instruction!")
	}
}

func main() {
	var vm = NewVM()

	for ok := true; ok; ok = vm.running {
		vm.eval(vm.fetch())
	}
}
