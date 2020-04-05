package main

import "fmt"

func main(){
	var n,d,m,c int
	fmt.Scan(&n)

	a:= make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&a[i])
	}
	fmt.Scan(&d,&m)

	for i:=0;i<n;i++{
		length,sum := 0,0
		for j:=i;j<n;j++{
			length++
			sum += a[j]
			if length == m && sum == d {
				c++
				break;
			}
			if sum > d || length > m{
				break
			}
		}
		// break brings here
	}
	fmt.Println(c)
}