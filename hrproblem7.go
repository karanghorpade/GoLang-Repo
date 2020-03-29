package main

import (
	"fmt"
	"sort"

)
func main() {
	var n int

	//var b,counter int
	fmt.Scan(&n)

	s:=set(n)
	c:=candles(n,s)

	fmt.Print(c)

}//main

func set(n int)[]int{
	a :=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&a[i])
	}
	sort.Ints(a)
	return a
}

func candles(n int,s []int)int{
	var b,counter int
	for i:=0;i<n;i++{
		b=s[n-1]
	}
	for _,v:=range s{
		if(b==v){
			counter++
		}
	}
	return counter
}

