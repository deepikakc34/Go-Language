# Understanding the data transfer through channels
package  main

import (
		"fmt"
)

func main() {

var ch chan string

go setfunc(ch)
go prntfunc(ch)

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
