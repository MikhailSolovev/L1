package main

import (
	"fmt"
	"reflect"
)

type Human struct {
	Name string
	Age  int
}

func (h Human) PrintName() {
	fmt.Println("Name:", h.Name)
}

func (h *Human) SetAge(age int) {
	h.Age = age
}

type ActionRef struct {
	*Human
}

type Action struct {
	Human
}

func main() {
	// Встраивание ptr структуры Human
	fmt.Println("----------With *Human----------")
	pr := ActionRef{&Human{}}
	t := reflect.TypeOf(pr)
	fmt.Println(t, "has", t.NumField(), "fields:")
	for i := 0; i < t.NumField(); i++ {
		fmt.Print(" field#", i, ": ", t.Field(i).Name, "\n")
	}
	fmt.Println(t, "has", t.NumMethod(), "methods:")
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Print(" method#", i, ": ", t.Method(i).Name, "\n")
	}

	// Встраивание структуры Human
	fmt.Printf("\n----------With Human----------\n")
	p := Action{Human{}}
	t = reflect.TypeOf(p)
	fmt.Println(t, "has", t.NumField(), "fields:")
	for i := 0; i < t.NumField(); i++ {
		fmt.Print(" field#", i, ": ", t.Field(i).Name, "\n")
	}
	fmt.Println(t, "has", t.NumMethod(), "methods:")
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Print(" method#", i, ": ", t.Method(i).Name, "\n")
	}

	// Структура ActionRef встраивает метод PrintName(), так и SetAge()
	// В то время как Action - только PrintName()
	// Но это не значит, что из объекта Action() нет доступа к методу SetAge()
	// p.SetAge() работает так же, как и pr.SetAge()
	// Это связано с тем, что в golang поведение следующих строчек одинаково
	// (&object).SetAge() и (object).SetAge()
}
