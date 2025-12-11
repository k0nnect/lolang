package core

import (
	"shared/pkg/data"
)

var clt = handler(func(ctx *functionCtx) {
	right := ctx.Stack.Pop()
	left := ctx.Stack.Pop()

	lt := false

	if left.GetInt() < right.GetInt() {
		lt = true
	}

	v, _ := data.NewValue(lt)
	ctx.Stack.Push(v)
})
