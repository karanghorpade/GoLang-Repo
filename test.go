package main


import (
	"fmt"
)

func main() {
 	var n,sl,valley int
	var s string
 	fmt.Scan(&n)

 	fmt.Scan(&s)

 	for _,v:=range s{
		if(v=='U'){
			sl++
		}else{
			sl--
		}

		if(sl==0 && v=='U'){
			valley++
		}
	}
	fmt.Println(valley)

}

