package  main

import (
		"fmt"
		//"time"
)

func main() {

var ch chan string

go setfunc(ch)
go prntfunc(ch)
//time.Sleep(50*1e9)

}

func setfunc(ch chan string) {

ch <- "Deepika"
ch <- "Praneeth"

}

func prntfunc(ch chan string){

for {

	input := <- ch
	fmt.Println(input)
}

}