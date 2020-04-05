package main


import (
	"fmt"
)

func main() {
	var p,n int
	fmt.Scan(&n,&p)
	if((p/2)<(n/2-p/2)){
		fmt.Println(p/2)
	}else{
		fmt.Println(n/2-p/2)
	}

}

