package core

var br = handler(func(ctx *functionCtx) {
	instr := ctx.Function.Instructions[ctx.InstrPtr]
	token := instr.Operand.GetInt()

	ctx.InstrPtr = int(token)
})
