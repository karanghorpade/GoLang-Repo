package main

import (
	"fmt"
)

func main(){

	ano()
}

func ano(){

	fmt.Println("ano runs")

	m := "Karan"

	func(m string){

		fmt.Println(m)

	}(m)


}
