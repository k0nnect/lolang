package function

import (
	"shared/pkg/data"
	"shared/pkg/opcodes"
	"shared/pkg/types"
)

const CustomFunctionStart = 1000000

type Function struct {
	Name         string
	Token        int // Identifier of the function similar to mdtoken in MSIL
	Instructions map[int]Instruction
	Locals       []Local
	Arguments    []Argument
	ReturnType   types.TypeCode
}

type Instruction struct {
	OpCode  opcodes.OpCode
	Operand data.Value
}

type Argument struct {
	Index int
	Name  string
	Type  types.TypeCode
}

type Local struct {
	Index           int
	Name            string
	Type            types.Type
	HasInitialValue bool
	InitialValue    data.Value
}
