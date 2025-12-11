package core

import (
	"shared/pkg/data"
)

var cgt = handler(func(ctx *functionCtx) {
	right := ctx.Stack.Pop()
	left := ctx.Stack.Pop()

	gt := false

	if left.GetInt() > right.GetInt() {
		gt = true
	}

	v, _ := data.NewValue(gt)
	ctx.Stack.Push(v)
})
