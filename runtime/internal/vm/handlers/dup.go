package handlers

import (
	"runtime/internal/vm/function"
)

var dup = Handler(func(ctx function.Ctx) {
	ctx.Stack.Push(ctx.Stack.Peek())
})
