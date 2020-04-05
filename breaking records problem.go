package main

import (
	"fmt"

)
func main() {
	var n,min,max,cmax,cmin int
	fmt.Scan(&n)
	a:=set(n)

	min = a[0]
	max = a[0]

	for i:=1;i<len(a);i++{
		if(a[i]>max){
			max = a[i]
			cmax++
		}else if(a[i]<min){
			min = a[i]
			cmin++
		}
	}
	fmt.Println(cmax,cmin)
}

func set(n int)[]int{
	a:=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&a[i])
	}
	return a
}

