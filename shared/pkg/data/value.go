package data

import (
	"shared/pkg/types"
)

type Value struct {
	Type types.TypeCode
	Data []byte
}

func (v *Value) GetInt() int64 {
	return int64(v.Data[0])
}
