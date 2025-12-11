package core

var bgt = handler(func(ctx *functionCtx) {
	right := ctx.Stack.Pop()
	left := ctx.Stack.Pop()

	if left.GetInt() <= right.GetInt() {
		return
	}

	instr := ctx.Function.Instructions[ctx.InstrPtr]
	token := int(instr.Operand.GetInt())

	ctx.InstrPtr = token
})
