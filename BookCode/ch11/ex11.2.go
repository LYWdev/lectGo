package main

import (
	"fmt"
	"time"
)

func main() {
	i :=1
	for {
		time.Sleep(time.Second)
		fmt.Println(i)
		i++

		switch i {
		case 2:
			fmt.Println("i값은 ",i, "이다.")
			i = 0
		}
	}
}