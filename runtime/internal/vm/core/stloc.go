package core

var stLoc = Handler(func(ctx *FunctionCtx) {
	instr := ctx.Function.Instructions[ctx.InstrPtr]
	token := instr.Operand.GetInt()

	ctx.Locals[int(token)] = ctx.Stack.Pop()
})
