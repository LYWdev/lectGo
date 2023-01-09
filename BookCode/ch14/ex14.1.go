package main

import "fmt"

func main()	{
	var a int = 500
	var p *int

	var b int = 10
	var c int = 20

	var p1 *int = &b
	var p2 *int = &b
	var p3 *int = &c

	p=&a
	fmt.Printf("P : %p\n",p)
	fmt.Printf("*P : %d\n",*p)
	*p = 100 // *P가 가리키는 메모리 공간의 값을 변경



	fmt.Printf("a의 값 : %d\n",a)
	fmt.Printf("p1 == p2 : %v\n",p1==p2)
	fmt.Printf("p2 == p3 : %v\n",p2==p3)

}
