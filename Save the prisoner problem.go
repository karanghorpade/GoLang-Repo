package main

import (
	"fmt"
)

func main(){
	var t int
	fmt.Scan(&t)
	for i:=0;i<t;i++{
		calculate()
	}

}

func calculate(){
	var n,m,s int
	fmt.Scan(&n,&m,&s)
	if((s-1+m)%n==0){
		fmt.Println(n)
	}else{
		fmt.Println((s-1+m)%n)
	}
}