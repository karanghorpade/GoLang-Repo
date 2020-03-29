package main
import(
"fmt"
)

func main(){
	var s,t,a,b,m,n,counter1,counter2 int

	fmt.Println("Enter S and t:")
	fmt.Scanln(&s,&t)

	fmt.Println("Enter a and b:")
	fmt.Scanln(&a,&b)

	fmt.Println("Enter m and n:")
	fmt.Scanln(&m,&n)
	arr1:=make([]int,m)
	for i:=0;i<m;i++{
		fmt.Scan(&arr1[i])
	}
	arr2:=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&arr2[i])
	}

	for _,v:=range arr1{
		if(s<=(v+a) && (v+a)<=t){
			counter1++
		}
	}
	for _,v:=range arr2{
		if(s<=(v+b) && (v+b)<=t){
			counter2++
		}
	}
	fmt.Println(counter1)
	fmt.Println(counter2)




}//main



