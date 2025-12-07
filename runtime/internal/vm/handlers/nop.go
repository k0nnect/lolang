package handlers

import (
	"runtime/internal/vm/function"
)

var nop = Handler(func(ctx function.Ctx) {
	// Do nothing
})
