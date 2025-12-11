package core

import "shared/pkg/data"

var cle = handler(func(ctx *functionCtx) {
	v1 := ctx.Stack.Pop()
	v2 := ctx.Stack.Pop()

	le := false

	if v1.GetInt() <= v2.GetInt() {
		le = true
	}

	v, _ := data.NewValue(le)
	ctx.Stack.Push(v)
})
