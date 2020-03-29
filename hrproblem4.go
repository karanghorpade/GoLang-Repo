package main

import (
	"fmt"
	//"math"
)
func main() {

	var principal, secondary int
	var row, col int
	fmt.Println("Enter number of rows:")
	fmt.Scanln(&row)
	fmt.Println("Enter number of columns:")
	fmt.Scanln(&col)

	fmt.Println("Enter Matrix:")
	matrix:=set(row,col)

	fmt.Println("Matrix:")
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf(" %d ",matrix[i][j])
			if(j==col-1){
				fmt.Println("")
			}
		}
	}

	fmt.Println("Sum of diagonal elements:")
	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
			if(i==j){
				principal+=matrix[i][j]
			}
			if(i+j==row-1){
				secondary+=matrix[i][j]
			}
		}
	}

	fmt.Println("Pricipal:",principal)
	fmt.Println("Secondary:",secondary)

	absolute_val:=abs(principal,secondary)
	fmt.Println("Absolute value is:",absolute_val)

}//main

func set(row,col int)[100][100]int{
	var matrix [100][100]int
	//m = make([][]int,row,col)
	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
			fmt.Println("Enter matrix element:")
			fmt.Scanln(&matrix[i][j])
		}
	}
	return matrix

}

func abs(principal,secondary int)int{
	abs:=principal-secondary
	if(abs<0) {
		abs= -abs
	}
	return abs
}

