package core

var pop = handler(func(ctx *functionCtx) {
	ctx.Stack.Pop()
})
