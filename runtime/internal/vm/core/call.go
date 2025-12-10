package core

var call = handler(func(ctx *functionCtx) {
	instr := ctx.Function.Instructions[ctx.InstrPtr]
	token := instr.Operand.GetInt()

	target := ctx.Vm.Functions[int(token)]
	fn := newFunctionCtx(ctx.Vm, &target)

	ctx.Vm.CallStack.Push(fn)

	for range target.Arguments {
		fn.Arguments = append(fn.Arguments, ctx.Stack.Pop())
	}

	fn.execute()
})
