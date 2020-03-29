package main
import(
	"fmt"
)

func main(){
	var x1,v1,x2,v2,diff int
	fmt.Scan(&x1,&v1,&x2,&v2)
	if(v1>v2){
		diff = (x1-x2)%(v2-v1)
		if(diff==0){
			fmt.Println("YES")
		}else{
			fmt.Println("NO")
		}

	}else{
		fmt.Println("NO")
	}


}//main
