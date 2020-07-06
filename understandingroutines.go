package main

import (
		"fmt"
)

func modif() {
	fmt.Println("in first go routine function")

	fmt.Println("ending first go routine")
}

func prtn() {
	fmt.Println("in second go routine function")


	fmt.Println("ending second go routine function")
}


func main(){

	fmt.Println("in main function")
	go modif()
	go prtn()
	fmt.Println("ending main function")
	
}
