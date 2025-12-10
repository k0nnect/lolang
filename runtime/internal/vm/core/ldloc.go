package core

var ldLoc = handler(func(ctx *functionCtx) {
	instr := ctx.Function.Instructions[ctx.InstrPtr]
	token := instr.Operand.GetInt()

	ctx.Stack.Push(ctx.Locals[int(token)])
})
