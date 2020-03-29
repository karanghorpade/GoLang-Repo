package main
import(
	"fmt"
)

func main(){

}//main
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
/*func set(n int)[]int{
	a:=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scanln(&a[i])
	}
	return a
}*/

