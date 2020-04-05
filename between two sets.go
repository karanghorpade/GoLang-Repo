package main

import "fmt"

func main(){
	var n,m,c int
	fmt.Scan(&n,&m)

	a := make([]int,n)
	b := make([]int,m)

	for i :=0;i<n;i++{
		fmt.Scan(&a[i])
	}
	for i:=0;i<m;i++{
		fmt.Scan(&b[i])
	}

	for i:=1;i<=100;i++{
		factor := true

		//check first
		for j:=0;j<n;j++{
			if i % a[j] != 0 {
				factor = false
				break
			}
		}
		// check second
		if factor {
			for j :=0;j<m;j++{
				if b[j]%i !=0 {
					factor = false
					break
				}
			}
		}
		//update counter
		if factor{
			c++
		}

	}
	fmt.Println(c)
}