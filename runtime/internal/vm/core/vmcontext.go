package core

import (
	"shared/pkg/data"
	"shared/pkg/function"
	"shared/pkg/vm"
)

const nonStdFunctionStart = 1000000

type vCtx struct {
	EntryPoint *function.Function
	Functions  map[int]function.Function
	CallStack  data.Stack[*functionCtx]
}

type functionCtx struct {
	Vm        *vCtx
	InstrPtr  int
	Function  *function.Function
	Stack     data.Stack[data.Value]
	Locals    map[int]data.Value
	Arguments []data.Value
	Running   bool
}

func RunVm(vm *vm.Vm) {
	ctx := newVmCtx(vm)
	ctx.execute()
}

func newVmCtx(vm *vm.Vm) *vCtx {
	ctx := &vCtx{
		Functions: vm.Functions,
	}

	ep := ctx.Functions[vm.EntryPoint]
	ctx.EntryPoint = &ep

	return ctx
}

// error halts the virtual machine and kills the running application
func (vm *vCtx) error(err error) {
	panic(err)
}

func (vm *vCtx) execute() {
	ctx := newFunctionCtx(vm, vm.EntryPoint)
	ctx.execute()
}

func newFunctionCtx(vm *vCtx, fn *function.Function) *functionCtx {
	var locals map[int]data.Value
	for _, local := range fn.Locals {
		if !local.HasInitialValue {
			continue
		}

		locals[local.Index] = local.InitialValue
	}

	return &functionCtx{
		Vm:       vm,
		InstrPtr: 0,
		Function: fn,
		Stack:    data.Stack[data.Value]{},
		Locals:   locals,
	}
}

func (fn *functionCtx) execute() {
	fn.Running = true

	for fn.Running {
		ptr := fn.InstrPtr

		instr := fn.Function.Instructions[fn.InstrPtr]

		// execute the instruction
		handlers[instr.OpCode](fn)

		// If the ptr has been changed than a branch instruction was called and changed the flow
		if ptr != fn.InstrPtr {
			continue
		}

		fn.InstrPtr++
	}
}
