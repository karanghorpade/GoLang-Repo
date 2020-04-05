package main

import "fmt"

func main(){

	var n,t int

	fmt.Scan(&n)

	a := make([]int,5)

	for i:=0;i<n;i++{
		fmt.Scan(&t)
		t = t -1
		//fmt.Println(t)
		a[t]++
	}
	//fmt.Println(a)
	max,occ := a[0],0

	for i,num := range a{
		if num > max{
			max = num
			occ = i+1
		}
	}
	fmt.Println(occ)
}
