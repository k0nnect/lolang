package data

import (
	"encoding/binary"
	"errors"
	"shared/pkg/types"
)

type Value struct {
	Type types.TypeCode
	Data []byte
}

func NewValue(v any) (Value, error) {
	switch x := v.(type) {
	case int64:
		data := make([]byte, 8)
		binary.BigEndian.PutUint64(data, uint64(x))
		return Value{
			Type: types.LoInt,
			Data: data,
		}, nil
	}

	return Value{}, errors.New("invalid value type")
}

func (v *Value) GetInt() int64 {
	return int64(v.Data[0])
}
