package mainpackage main
import(
"fmt"
)

func main(){
	var n int
	fmt.Scanln(&n)
	a:=set(n)
	grade(a)

}//main

func set(n int)[]int{
	a:=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scanln(&a[i])
	}
	return a
}

func grade(a[]int){
	for _,v:=range a{
		if (v>=38){
			if((v+(5-(v%5))-v)<3){
				v = v+(5-(v%5))
			}

		}
		fmt.Println(v)
	}
}
