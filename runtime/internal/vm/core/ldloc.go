package core

var ldLoc = Handler(func(ctx *FunctionCtx) {
	instr := ctx.Function.Instructions[ctx.InstrPtr]
	token := instr.Operand.GetInt()

	ctx.Stack.Push(ctx.Locals[int(token)])
})
