package main

import (
	"fmt"
)

func main() {
	for i:=2; i<10; i++{
		switch i {
		case 3:
			continue
		case 4:
			continue
		case 5:
			continue
		case 6:
			continue
		}
		for j:=1; j<10; j++{
			fmt.Println(i,"*", j,"=",i*j)
		}
	}
}