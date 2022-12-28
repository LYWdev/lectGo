package main

import "fmt"

func main() {

	//오버플로우 예시
	var x int8 = 127

	fmt.Printf("%d < %d +1 : %v",x, x, x <x+1 )
	fmt.Printf("x\t= %4d, %08b\n",	x,	x)
	fmt.Printf("x + 1 \t=%4d, %08b\n",	x+1,	x+1 )
	fmt.Printf("x + 2 \t=%4d, %08b\n",	x+2,	x+2 )
	fmt.Printf("x + 3 \t=%4d, %08b\n",	x+3,	x+3 )

	//언더플로우 예시
	var y int8 = -128

	fmt.Printf("%d > %d - 1: %v\n ",	y,	y,	y>y-1)
	fmt.Printf("y\t=%4d, %08b\n ",	y,	y)	
	fmt.Printf("y-1\t=%4d, %08b\n ",	y-1,	y-1)

	}
	//오퍼플로우 값은 모두 버려짐
