package main

import (
	"fmt"
	"vbi-cosmos-basic/example"
)

func Execution() {
	person := &example.Person{
		Name: "Thang",
		Age:  1,
	}

	fmt.Println(person)
}
