package vm

import (
	"shared/pkg/function"
)

type VCtx struct {
	Functions map[int]function.Function
}

// Error halts the virtual machine and kills the running application
func (vm *VCtx) Error(err error) {
	panic(err)
}
