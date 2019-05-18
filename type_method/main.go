package main

import (
	"fmt"
	"github.com/inagacky/go_practice/type_method/method"
)

func main() {

	s := method.MyString("Fuga")
	s.AddHoge()
	fmt.Println(s.AddHoge())
}