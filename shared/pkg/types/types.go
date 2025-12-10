package types

const (
	LoInt = TypeCode(iota)
	LoVoid
	LoString
	LoBool
)

var Types = []Type{
	{
		Name: "int",
		Code: LoInt,
		Size: 8,
	},
	{
		Name: "bool",
		Code: LoBool,
		Size: 1,
	},
	{
		Name: "void",
		Code: LoVoid,
		Size: 0,
	},
	{
		Name: "string",
		Code: LoString,
	},
}

func GetTypeByName(name string) *Type {
	for _, t := range Types {
		if t.Name == name {
			return &t
		}
	}

	return nil
}

func GetTypeByCode(code TypeCode) *Type {
	for _, t := range Types {
		if t.Code == code {
			return &t
		}
	}

	return nil
}
