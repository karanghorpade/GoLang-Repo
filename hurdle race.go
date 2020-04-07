package main

import (
	"fmt"
)

func main() {
	var hurdle,height int

	fmt.Scan(&hurdle,&height)
	a:=make([]int,hurdle)
	for i:=0;i<hurdle;i++{
		fmt.Scan(&a[i])
	}
	highest := 0
	for i:=0;i<hurdle;i++{
		if(a[i]>highest){
			highest=a[i]
		}
	}
	fmt.Println(highest)
	if(height>highest){
		fmt.Println(0)
	}else{
		n:=highest-height
		fmt.Println(n)
	}

}
