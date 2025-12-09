package data

import "testing"

func TestNewValue(t *testing.T) {
	num := int64(5)

	v, _ := NewValue(num)
	num1 := v.GetInt()

	if num != num1 {
		t.Errorf("Num: %d is not equal to Value: %d", num, num1)
	}
}
