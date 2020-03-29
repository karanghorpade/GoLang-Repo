package main

import (
"fmt"
)

func main() {

	for i := 0; i <= 10; i++ {

		c := make(chan int, 1)

		c <- 12

		fmt.Println(<-c)
	}
}
