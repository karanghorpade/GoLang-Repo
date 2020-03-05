package main

import "fmt"

func main(){

	x:=[]string{"james","bond","shaken","not stirred"}
	y:=[]string{"karan","vijay","ghorpade","homelander","translucent"}

	z:=[][]string{x,y}

	for i,v:=range z{
		fmt.Println(i,v)

	}


}
