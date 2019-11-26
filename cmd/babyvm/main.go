package main

// All available instructions for our VM
const (
	PSH = iota
	ADD
	POP
	SET
	HLT
)

var program = [...]uint8{
	PSH, 5,
	PSH, 6,
	ADD,
	POP,
	HLT}

type VM struct {
	running bool
	pc      uint32
}

func NewVM() *VM {
	return &VM{
		running: false,
		pc:      0,
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
	}
}

func main() {
	var vm = NewVM()

	for ok := true; ok; ok = vm.running {
		vm.eval(vm.fetch())
	}
}
