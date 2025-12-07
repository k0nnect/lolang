package vm

import (
	"shared/pkg/function"
)

type VCtx struct {
	Functions map[int]function.Function
}
