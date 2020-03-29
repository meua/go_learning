package encap

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

// func (e *Employee) String() string {
// 	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
// }

func (e *Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	e3 := Employee{Id: "1", Name: "Mary", Age: 40}
	e2 := new(Employee) //返回指针
	e2.Id = "2"
	e2.Age = 22
	e2.Name = "Rose"
	e4 := &Employee{}
	e4.Id = "4"
	e4.Name = "cane"
	e4.Age = 50
	t.Log(e)
	t.Log(e1)
	t.Log(e3)
	t.Log(e1.Id)
	t.Log(e2)
	t.Log(e4)
	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)
	t.Logf("e3 is %T", &e3)
	t.Logf("e4 is %T", e4)
}

func TestStructOperations(t *testing.T) {
	//e := &Employee{"0", "Bob", 20}
	e := Employee{"0", "Bob", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
