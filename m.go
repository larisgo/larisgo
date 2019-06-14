package main

import (
	"fmt"
	"reflect"
)

type Testa interface {
	a()
}
type Testb interface {
	a()
}
type AAA struct {
	x bool
}
type ssAAA struct {
	x bool
}

func (this ssAAA) a() {

}
func (this *ssAAA) b() {

}
func main() {
	var a interface{}
	var b interface{}
	var c interface{}
	a = &AAA{x: false}
	b = &ssAAA{x: true}
	c = ""
	fmt.Println(a == b)
	fmt.Println(reflect.TypeOf(a) == reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(Testb(&ssAAA{x: true})).String())
	fmt.Println(reflect.TypeOf(b).String())
	fmt.Println(reflect.TypeOf(a).Kind() == reflect.TypeOf(b).Kind())
	fmt.Println(a == c)
	fmt.Println(Testb(&ssAAA{x: true}))

}
