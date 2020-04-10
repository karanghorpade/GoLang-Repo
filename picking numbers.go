package main


import (
"fmt"
"math"
)

func main() {
	var result float64
	var n int
	var arr[101]float64
	fmt.Scan(&n)
	a:=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&a[i])
	}

	for i:=0;i<len(a);i++{
		index:=a[i]
		arr[index]++
	}

	for i:=1;i<100;i++{
		result=math.Max(result,(arr[i]+arr[i-1]))
	}

	fmt.Println(result)


}




