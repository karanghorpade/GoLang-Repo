package main

import (
	"fmt"
	//"sort"
)

func main(){
	var d,m,y int
	fmt.Scan(&y)

	if(y<=1917){
		if(y%4==0){
			d=256-244
			m=9
			fmt.Print(d,".0",m,".",y)
		}else{
			d=256-243
			m=9
			fmt.Print(d,".0",m,".",y)
		}

	}else if(y==1918){
		m=9
		d=256-230
		fmt.Print(d,".0",m,".",y)
	}else if(1918<y && y<=2700){
		if((y%400==0)||(y%4==0&&y%100!=0)){
			d=256-244
			m=9
			fmt.Print(d,".0",m,".",y)
		}else{
			d=256-243
			m=9
			fmt.Print(d,".0",m,".",y)
		}
	}

}