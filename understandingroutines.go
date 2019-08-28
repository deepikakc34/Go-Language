package main

import (
		"fmt"
	//	"time"
)

func modif() {
	fmt.Println("in first go routine function")
	//time.Sleep(10 * 1e9)
	fmt.Println("ending first go routine")
}

func prtn() {
	fmt.Println("in second go routine function")
//	time.Sleep(10 * 1e9)

	fmt.Println("ending second go routine function")
}


func main(){

	fmt.Println("in main function")
	go modif()
	go prtn()
	//time.Sleep(10 * 1e9)
	fmt.Println("ending main function")
	
}