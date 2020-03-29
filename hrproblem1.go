package main

import (
	"fmt"
)

func main(){

	var size int
	fmt.Scan(&size)
	b:= set(size)

	add:=sum(b)
	fmt.Println("Sum of the array elements: ",add)
}

func set(size int)[]int{

	a := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&a[i])
	}
	return a
}

func sum(s []int)int{
	sum := 0
	for _,v:=range s{
		sum +=v
	}
	return sum
}
