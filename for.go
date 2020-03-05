package main

import "fmt"

func main(){

	i := 1
	for i<=3{

		fmt.Println("karan")
		i=i+1
	}

	for i=1; i<=5; i++{

		if i%2==0{
			continue
		}

		fmt.Println(i)
	}

}
