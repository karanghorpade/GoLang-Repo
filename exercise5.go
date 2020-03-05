package main

import(
	"fmt"
)

type(
	x = int
)

var y x = 10

func main(){
	
	//fmt.Printf("%T",y)

	y := string('y')
	fmt.Printf("%T",y)
	
	fmt.Println(y)
	
}