package main

import "fmt"

func main(){
	//fmt.Println("Enter size of array:")
	size := 3
	//fmt.Scan(&size)
	//fmt.Println("Enter first array:")
	s1:=set(size)
	//fmt.Println("Enter second array:")
	s2:=set(size)


	s3:=compare(s1,s2)
	for _,v:=range s3{
		fmt.Print(v)
		fmt.Print(" ")
	}


}
func set(size int)[]int{

	s := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&s[i])
	}
	return s
}

func compare(s1,s2[]int)[]int{
	s3:=[]int{}
	var counter1 int
	var counter2 int
	for i:=0;i<3;i++{
		if s1[i]>s2[i]{
			counter1 ++
		}else if s1[i]<s2[i]{
			counter2++
		}

	}
	s3=append(s3, counter1,counter2)
	return s3

}


