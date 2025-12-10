package core

import (
	"shared/pkg/data"
)

var cgt = handler(func(ctx *functionCtx) {
	v1 := ctx.Stack.Pop()
	v2 := ctx.Stack.Pop()

	gt := false

	if v1.GetInt() > v2.GetInt() {
		gt = true
	}

	v, _ := data.NewValue(gt)
	ctx.Stack.Push(v)
})
