package main

import "fmt"


func main() {

	var l int
	var arr[3][3] string
	var fg1, fg4 bool = true, true
	var temp[3] string

	arr[0][0] = "X"
	arr[0][1] = "X"
	arr[0][2] = "O"
	arr[1][0] = "O"
	arr[1][1] = "X"
	arr[1][2] = "X"
	arr[2][0] = "O"
	arr[2][1] = "O"
	arr[2][2] = "X"

	l=len(arr)

	for i:=0;i<l;i++ {
		var fg2, fg3 bool = true, true
		for j:=0;j<l;j++ {
	// TO CHECK IF GAME EXISTS IN DIAGONAL
			 if i==j {
				 if i < l-1 && arr[i][j] != arr[i+1][j+1] {
					 fg1 = false
					}
					 if i==l-2 && fg1==true {
						fmt.Printf("%s is the winner of the game", arr[i][j])
						fmt.Println(" ")
						goto label
					 }
				  
			} 
	// TO CHECK IF GAMES EXISTS IN THE ROWS	
			if  j<l-1 && arr[i][j] != arr[i][j+1]{
				fg2 = false
			}
				if j==l-2 && fg2==true {
					fmt.Printf("%s is the winner of the game", arr[i][j])
					fmt.Println(" ")
					goto label

				}
				
				
	// TO CHECK IF GAME EXISTS IN THE COLUMNS
			 
			if  j<l-1 && arr[j][i] != arr[j+1][i] {
				fg3 = false
			}
				if j==l-2 && fg3==true {
				fmt.Printf("%s is the winner of the game", arr[i][j])
				fmt.Println(" ")
				goto label
				}
	// TO CHECK IF THE GAME EXISTS IN THE DIAGONAL	
				
			if j==l-1-i {
				temp[i]=arr[i][j]
				if i>0 && temp[i] != temp[i-1] {
					fg4 = false
				}
			}
				if i==l-1 && fg4==true {
					fmt.Printf("%s is the winner of the game", temp[i])
					fmt.Println(" ")
					goto label

				} 
		
			
	


			} 
	
		 }
	// NO WIN CONDITION SEEN IN THE GAME	 
	fmt.Println("The game is a draw")
	label: return
	
}
