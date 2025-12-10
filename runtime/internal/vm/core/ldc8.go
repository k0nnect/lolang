package core

var ldc8 = handler(func(ctx *functionCtx) {
	value := ctx.Function.Instructions[ctx.InstrPtr].Operand
	ctx.Stack.Push(value)
})
