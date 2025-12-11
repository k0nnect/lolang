package core

import "shared/pkg/data"

var bt = handler(func(ctx *functionCtx) {
	v1 := ctx.Stack.Pop()

	tr := data.MustNewValue(true)

	if !v1.Equal(&tr) {
		return
	}

	instr := ctx.Function.Instructions[ctx.InstrPtr]
	token := int(instr.Operand.GetInt())

	ctx.InstrPtr = token
})
