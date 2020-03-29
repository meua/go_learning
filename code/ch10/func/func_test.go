package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
	//tsSF := timeSpent(slowFun)
	//t.Log(tsSF(10))

	t.Log(timeSpent(slowFun)(10))
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear resources.")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	//t.Log("start")
	panic("err")
	t.Log("not print")
}

func TestDefer1(t *testing.T) {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\n", i)
	}
}

func Defer2() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func TestDefer2(t *testing.T) {
	t.Logf("%d\n", Defer2())
}
