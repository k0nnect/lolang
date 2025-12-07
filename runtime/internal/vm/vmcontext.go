package vm

import (
	vmFunc "runtime/internal/vm/function"
	"shared/pkg/data"
	"shared/pkg/function"
)

type VCtx struct {
	EntryPoint *function.Function
	Functions  map[int]function.Function
	CallStack  data.Stack[*vmFunc.Ctx]
}

// Error halts the virtual machine and kills the running application
func (vm *VCtx) Error(err error) {
	panic(err)
}
