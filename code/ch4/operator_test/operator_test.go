package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 2, 4}
	//	c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)
}

// 解释：&^ ,符号左边，0保持不变，1变为相反
func TestBitClear(t *testing.T) {
	a := 7            //0111
	a = a &^ Readable // 0111 &^ 0001 -> 0110
	t.Log(Readable)
	t.Log(a)
	a = a &^ Executable // 0110 &^ 0100 -> 0010
	t.Log(Executable)
	t.Log(a)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
