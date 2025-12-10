package core

var dup = handler(func(ctx *functionCtx) {
	ctx.Stack.Push(ctx.Stack.Peek())
})
