package main

import (
	"fmt"
	"sort"
	//"math"
)
func main() {
	//var min,max int
	a:=make([]int,5)
	for i:=0;i<len(a);i++{
		fmt.Scan(&a[i])
	}
	//fmt.Scan()
	sort.Ints(a)

	mi:=min(a)
	ma:=max(a)
	fmt.Println(mi,ma)


}//main

func min(a[]int)int{
	var min int
	for i:=0;i<len(a)-1;i++{
		min += a[i]
	}
	return min

}

func max(a[]int)int{
	var max int
	for i:=1;i<len(a);i++{
		max += a[i]
	}
	return max
}
