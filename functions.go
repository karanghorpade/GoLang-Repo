package main

import(

	"fmt"
)

func main(){

	name:=random("karan")
	fmt.Println("The name is ", name)
}

func random(s string) string{

	i:=0
	for i=0;i<=10;i++{
		fmt.Println(s)
	}

	return s

}
