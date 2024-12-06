/*
	Разработать программу, которая в рантайме способна
	определить тип переменной: int, string, bool, channel
	из переменной типа interface{}.
*/

package main

import (
	"fmt"
	"reflect"
)

type MyType interface{}

func main() {

	i := 31337
	s := "hello"
	b := true
	c := make(chan int)

	fmt.Println(getTypeFmt(i))
	fmt.Println(getTypeReflect(i))
	fmt.Println(getTypeSwitch(i))

	fmt.Println(getTypeFmt(s))
	fmt.Println(getTypeReflect(s))
	fmt.Println(getTypeSwitch(s))

	fmt.Println(getTypeFmt(b))
	fmt.Println(getTypeReflect(b))
	fmt.Println(getTypeSwitch(b))

	fmt.Println(getTypeFmt(c))
	fmt.Println(getTypeReflect(c))
	fmt.Println(getTypeSwitch(c))

}

func getTypeFmt(v MyType) string {
	return fmt.Sprintf("%T", v) // на самом деле здесь под капотом тоже используется пакет reflect
}

func getTypeReflect(v MyType) string {
	return reflect.TypeOf(v).String()
}

func getTypeSwitch(v MyType) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	default:
		return "unknown"
	}
}
