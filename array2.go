package main

import "fmt"

func main(){

	x:=[]int{42,43,44,45,46,47,48,49,50,51}


	z:=make([]int,10,100)

	x = append(x,6,5,67,7)
	fmt.Println(x)

	y:=[]int{1,2,3,4}

	x=append(x, y...)
	fmt.Println(x)

}
