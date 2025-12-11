package core

import "shared/pkg/data"

var cge = handler(func(ctx *functionCtx) {
	v1 := ctx.Stack.Pop()
	v2 := ctx.Stack.Pop()

	ge := false

	if v1.GetInt() >= v2.GetInt() {
		ge = true
	}

	v, _ := data.NewValue(ge)
	ctx.Stack.Push(v)
})
