package core

import "shared/pkg/data"

var cle = handler(func(ctx *functionCtx) {
	right := ctx.Stack.Pop()
	left := ctx.Stack.Pop()

	le := false

	if left.GetInt() <= right.GetInt() {
		le = true
	}

	v, _ := data.NewValue(le)
	ctx.Stack.Push(v)
})
