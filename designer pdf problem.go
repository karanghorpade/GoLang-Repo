package main
//designer pdf viewer
import (
	"fmt"
)

func main() {
	h,word:=set()
	area:=calculate(h,word)
	fmt.Println(area)

}

func set() ([26]int,[]byte) {

	var h[26]int
	for i:=0;i<26;i++{
		fmt.Scan(&h[i])
	}
	var word []byte
	fmt.Scan(&word)

	return h,word
}

func calculate(h[26]int ,word[]byte )int{
	var max int
	for i:=0;i<len(word);i++{
		a:=word[i]
		height:=h[a-97]
		if(height>max){
			max=height
		}
	}
	var area int
	area=max*len(word)
	return area
}
