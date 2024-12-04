package main

import "fmt"

type Human struct {
	fname string
	lname string
	age   int
	sex   string
}

func (h Human) GetFullName() string {
	return fmt.Sprintf("%s %s", h.fname, h.lname)
}

func (h Human) GetAge() int {
	return h.age
}

func (h Human) GetSex() string {
	return h.sex
}

// в Go нет наследования в традиционном его понимании
// наследование реализовано через встраивание (композицию)
type Action struct {
	mname string
	Human
}

// мы не можем перепоределять методы встроенной структуры
// но можем затенять их (shadowing)
// метод GetFullName() в Action будет затенять метод GetFullName() в Human
func (a Action) GetFullName() string {
	return fmt.Sprintf("%s %s %s", a.fname, a.mname, a.lname)
}

func main() {
	fm := Human{"Farrokh", "Bulsara", 45, "male"}
	a := Action{"Bomi", fm}

	fmt.Println(a.GetFullName())
	fmt.Println(a.GetAge())
	fmt.Println(a.GetSex())

	// мы все ещё можем обратиться к затенённому методу
	// встроенной структуры с помощью такой конструкции
	fmt.Println(a.Human.GetFullName())
}
