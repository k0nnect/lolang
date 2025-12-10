package core

import (
	"shared/pkg/data"
)

var clt = handler(func(ctx *functionCtx) {
	v1 := ctx.Stack.Pop()
	v2 := ctx.Stack.Pop()

	lt := false

	if v1.GetInt() < v2.GetInt() {
		lt = true
	}

	v, _ := data.NewValue(lt)
	ctx.Stack.Push(v)
})
