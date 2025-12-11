package core

import "shared/pkg/data"

var cge = handler(func(ctx *functionCtx) {
	right := ctx.Stack.Pop()
	left := ctx.Stack.Pop()

	ge := false

	if left.GetInt() >= right.GetInt() {
		ge = true
	}

	v, _ := data.NewValue(ge)
	ctx.Stack.Push(v)
})
