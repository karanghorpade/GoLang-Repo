package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	arr:=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scanln(&arr[i])
	}
	for i:=0;i<len(arr);i++{
		calculate(arr[i])
	}

}
func calculate(a int){
	var height int
	height=1
	for j:=1;j<=a;j++{
		if(j%2!=0){
			height=height*2
		}else{
			height=height+1
		}

	}
	fmt.Println(height)

}
