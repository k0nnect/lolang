package vm

import "shared/pkg/function"

type Vm struct {
	EntryPoint int
	Functions  map[int]function.Function
}
