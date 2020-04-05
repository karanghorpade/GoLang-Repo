package main


import (
	"fmt"
	"math"
)

func main() {
 	var q int

 	fmt.Scan(&q)
 	for i:=0;i<q;i++{
 		compute()
	}
}

func compute(){
	var x,y,z float64
	fmt.Scan(&x,&y,&z)
	if(math.Abs(z-y)<math.Abs(z-x)){
		fmt.Println("Cat B")
	}else if(math.Abs(z-y)>math.Abs(z-x)){
		fmt.Println("Cat A")
	}else{
		fmt.Println("Mouse C")
	}
}


