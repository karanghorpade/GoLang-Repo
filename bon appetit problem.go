package main

import (
	"fmt"
	//"sort"
)

func main(){

	var n,k,refund,bc,ba,total int
	fmt.Scan(&n,&k)
	a:=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&a[i])
	}
	fmt.Scan(&bc)

	for i:=0;i<n;i++{
		if(i!=k){
			total+=a[i]
		}
	}
	ba=total/2
	if(ba==bc){
		fmt.Println("Bon Apetite")
	}else{
		refund=bc-ba
		fmt.Println(refund)
	}


}