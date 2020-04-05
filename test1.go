package main

import (
	"fmt"
	"strings"
)

func main(){
	var n int
	fmt.Scan(&n)
	a:=make([]string,n)
	for i:=0;i<n;i++{
		fmt.Scan(&a[i])
	}
	s := strings.Join(a, "")
	fmt.Println(s)
}



