package main

import (
	"fmt"
)

func main() {
	var n, k, q, m, j int
	fmt.Scan(&n, &k, &q)
	k = k % n
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < q; i++ {
		fmt.Scanln(&m)
		j = m - k
		if (j < 0) {
			fmt.Println(a[n+j])
		} else {
			fmt.Println(a[j])
		}
	}

}