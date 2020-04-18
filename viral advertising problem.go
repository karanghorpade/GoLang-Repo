package main

import (
	"fmt"
	"math"
)

func main() {

	var n int
	fmt.Scan(&n)
	cumulative:=calculate(n)
	fmt.Println(cumulative)

}

func calculate(n int)int{
	var shared, liked, cumulative float64
	shared =5
	liked, cumulative =2,2

	for i:=1;i<n;i++{
		shared=liked*3
		liked=math.Floor(shared/2)
		cumulative=cumulative+liked
	}
	return int(cumulative)
}

