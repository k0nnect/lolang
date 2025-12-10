package core

var ldarg = handler(func(ctx *functionCtx) {
	idxOp := ctx.Function.Instructions[ctx.InstrPtr].Operand
	idx := idxOp.GetInt()

	ctx.Stack.Push(ctx.Arguments[idx])
})
