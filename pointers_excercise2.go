package main

import (
	"fmt"
)

type person struct {
	age int
}

func changeMe(p *person) int {

	(*p).age = 20

	return (*p).age

}

func main() {

	p1 := person{
		age: 15,
	}
	fmt.Println(p1.age)

	s := changeMe(&p1)

	fmt.Println(s)

	fmt.Println(p1.age)

}
