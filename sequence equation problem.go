package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a:=make([]int,n+1)
	for i:=1;i<=n;i++{
		var k int
		fmt.Scan(&k)
		a[k]=i
	}

	for i:=1;i<=n;i++{
		fmt.Println(a[a[i]])
	}

}
