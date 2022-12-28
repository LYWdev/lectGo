package main

import "fmt"

func main(){
	const ( 
		Red 	int = iota
		Blue 	int = iota
		Green 	int = iota
	)

	const ( 
		C1 	uint = iota + 5
		C2 	
		C3
	)

	const ( 
		BitFlag1 uint = 1 << iota
		BitFlag2
		BitFlag3
		BitFlag4
	)

	fmt.Println(Red,Green,Blue)
	fmt.Println(C1,C2,C3)
	fmt.Println(BitFlag1,BitFlag2,BitFlag3,BitFlag4)
}