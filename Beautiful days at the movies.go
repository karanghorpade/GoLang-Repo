package main

import (
	"fmt"
	"math"

)

func main() {
	var a,b,count int
	var k float64
	fmt.Scan(&a,&b,&k)
	arr,arr1:=reverse(a,b)

	for i:=0;i<len(arr);i++{
		q:=math.Abs(float64(arr[i]-arr1[i]))
		z:=q/k
		if(math.Mod(z,1)==0){
			count++
		}
	}
	fmt.Println(count)
}

func reverse(a,b int)([]int,[]int){
	var reminder int
	var arr[]int
	for i:=a;i<=b;i++{
		arr=append(arr,i)
	}
	var arr1[]int
	for _,v:=range arr {
		var reverse int
		for (v != 0) {
			reminder = (v % 10)
			reverse = (reverse * 10) + reminder
			v = v / 10

		}
		arr1=append(arr1,reverse)

	}
	return arr,arr1
}

