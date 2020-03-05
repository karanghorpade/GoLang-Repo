package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretagent struct {
	person
	ltk bool
}

func main() {

	sa1 := secretagent{

		person: person{
			"james",
			"bond",
		},

		ltk: true,
	}

	sa1.speak()

}

func (s secretagent) speak() {

	fmt.Println("It is", s.ltk,"that I am", s.first, s.last)

}
