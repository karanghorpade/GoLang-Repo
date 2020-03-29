package main

import (
	"fmt"
	//"math"
)
func main() {
	var b int
	var x,y,z float64
	var n,positive,negative,zeros int
	//fmt.Println("Enter numbers:")
	fmt.Scanln(&n)
	a:=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&a[i])
	}

	for i,v:=range a{
		i++
		b=+i
		if(v>0){
			positive++
		}else if(v<0){
			negative++
		}else if(v==0){
			zeros++
		}
	}

	//fmt.Println("b:",b)
	x = float64(positive)/float64(b)
	y = float64(negative)/float64(b)
	z=float64(zeros)/float64(b)
	fmt.Println("%.2f",x)
	fmt.Println("%.2f",y)
	fmt.Println("%.2f",z)




}//main




