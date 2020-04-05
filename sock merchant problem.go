package main


import (
	"fmt"
	"sort"
)

func main() {

	var n,c, counter int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	sort.Ints(arr)

	for i := 0; i < n; i++ {
		c=1
		if(arr[i]!=0){
			for j:=i+1;j<n;j++{
				if(arr[i]==arr[j]){
					c++
					arr[j]=0
				}
			}
			counter+=c/2;
		}

	}
	fmt.Println(counter)
}