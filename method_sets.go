package main

import(
	"fmt"
)

type person struct{

	name string

}

type human interface{

	speak() string
}

func (p *person)speak() string{

	return p.name
}

func saySomething(h human){

	fmt.Println(h.speak())
}

func main(){

	a:=person{s:"jack"}
	saySomething(&a)
}