package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i:=0;i<t;i++{
		calculate()
	}
}
func calculate(){
	var n,r,count int
	fmt.Scan(&n)
	r = n
	for(r > 0){
		m := r%10
		r = r / 10
		if(m > 0 && n % m == 0) {
			count++
		}

	}
	fmt.Println(count)
}
