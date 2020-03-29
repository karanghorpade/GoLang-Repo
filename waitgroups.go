package main

import(

	"fmt"
	//"time"
	"sync"

)

func main(){

	var wg sync.WaitGroup
	wg.Add(2)
	go func(){

		fmt.Println("First async function")
		wg.Done()
	}()

	go func(){

		fmt.Println("Second async function")
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("done")

}