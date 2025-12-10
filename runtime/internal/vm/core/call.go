package core

import (
	"runtime/pkg/std"
	"shared/pkg/data"
	"shared/pkg/types"
)

var call = handler(func(ctx *functionCtx) {
	instr := ctx.Function.Instructions[ctx.InstrPtr]
	token := instr.Operand.GetInt()

	// Check to see if this should be handled by the std lib
	if token < CustomFunctionStart {
		stdFn := std.StdLib[int(token)]

		var args []data.Value

		for range stdFn.Arguments {
			args = append(args, ctx.Stack.Pop())
		}

		ret := stdFn.Execute(args)

		if ret.Type != types.LoVoid {
			ctx.Stack.Push(ret)
		}

		return
	}

	target := ctx.Vm.Functions[int(token)]
	fn := newFunctionCtx(ctx.Vm, &target)

	ctx.Vm.CallStack.Push(fn)

	for range target.Arguments {
		fn.Arguments = append(fn.Arguments, ctx.Stack.Pop())
	}

	fn.execute()
})
