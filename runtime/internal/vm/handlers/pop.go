package handlers

import (
	"runtime/internal/vm/function"
)

var pop = Handler(func(ctx function.Ctx) {
	ctx.Stack.Pop()
})
