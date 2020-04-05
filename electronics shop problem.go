package main


import (
	"fmt"
)

func main() {
	var budget, keyboards, usb int
	var temp int
	fmt.Scan(&budget,&keyboards,&usb)
	k:=set(keyboards)
	pd:=set(usb)

	for i:=0;i<keyboards;i++{
		for j:=0;j<usb;j++{
			if((k[i]+pd[j])<=budget){
				if((k[i]+pd[j])>temp){
					temp=k[i]+pd[j]
				}
			}
		}
	}
	if(temp==0){
		fmt.Println(-1)
	}else{
		fmt.Println(temp)
	}


}

func set(size int)[]int{
	arr:=make([]int,size)
	for i:=0;i<size;i++{
		fmt.Scan(&arr[i])
	}
	return arr
}
func show(a []int){
	for _,v:=range a{
		fmt.Print(",",v)
	}
}


