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

func TestBitClear(t *testing.T) {
	a := 7 //0111
	a = a &^ Readable
	a = a &^ Executable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}

/*go语言中按位取反写法是^, 所以 a&^b 其实是 a&(^b) 利用运算符优先级省略掉括号的写法而已.
下面的测试方法可以自行验证一下. 其它语言中关键字不同写法可能是 a&(~b), 但算法是一样的.*/
func TestBitClear1(t *testing.T) {
	flag := Readable | Executable
	t.Log(flag&^Readable == Executable)
	t.Log(flag&^Executable == Readable)
	t.Log(flag&(^Readable) == Executable)
	t.Log(flag&(^Executable) == Readable)
}
