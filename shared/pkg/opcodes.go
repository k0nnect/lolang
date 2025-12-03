package pkg

type OpCode int

const (
	// Arithmetic Operations
	Add OpCode = iota
	Sub
	Mul
	Div
	Rem
	Xor
	Not
	Or
	Neg

	// Branch Operations
	Br
	Be
	Bne
	Blt
	Bgt

	// Equality
	Cmp
	Clt
	Cgt

	// Stack Control
	Ldc4
	Pop

	// Misc
	Nop
	Ret
)
