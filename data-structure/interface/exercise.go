package main

import "fmt"

type speaker interface {
	speak()
}

type english struct {
	language string
}

type hindi struct {
	language string
}

func (e english) speak() {
	fmt.Println(e.language)
}

func (h hindi) speak() {
	fmt.Println(h.language)
}

func sayHello(s speaker) {
	s.speak()
}

func main() {
	h := hindi{"Namaste"}
	e := english{"Hello"}
	sayHello(h)
	sayHello(e)
}
