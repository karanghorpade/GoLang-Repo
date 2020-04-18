package main

import (
	"fmt"
)

func main() {
	var t int
	var b bool
	fmt.Scanln(&t)
	for i:=0;i<t;i++{
		b=set()
		if(b==true){
			fmt.Println("YES")
		}else{
			fmt.Println("NO")
		}
	}
}

func set()bool{
	var n, k,count int
	fmt.Scan(&n,&k)
	a:=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&a[i])
		if(a[i]<=0){
			count++
		}
	}

	if(count==k){
		return false
	}else{
		return true
	}
}