package core

var stLoc = handler(func(ctx *functionCtx) {
	instr := ctx.Function.Instructions[ctx.InstrPtr]
	token := instr.Operand.GetInt()

	ctx.Locals[int(token)] = ctx.Stack.Pop()
})
