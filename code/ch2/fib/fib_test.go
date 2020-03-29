package fib

import (
	"math"
	"testing"
)

var a int

func TestFibList(t *testing.T) {
	// var a int = 1
	// var b int = 1
	// var (
	// 	a int = 1
	// 	b     = 1
	// )
	//a := 1
	// a := 1
	//b := 1
	a, b := 1, 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}

}

func TestExchange(t *testing.T) {
	a := math.MaxInt32
	b := math.MinInt32
	// tmp := a
	// a = b
	// b = tmp
	a, b = b, a
	t.Log(a, b)
}
