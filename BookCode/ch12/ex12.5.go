package main

import "fmt"

func main() {
	a:=[5]int{1,2,3,4,5}
	b:=[5]int{500,400,300,200,100}

	for i, v:=range a{
		fmt.Printf("a[%d]=%d\n",i,v)
	}

	for i, v:=range b{
		fmt.Printf("b[%d]=%d\n",i,v)
	}

	// 배열의 타입과 크기는 모두 일치해야한다.
	b=a
	fmt.Println("변수 복사 이후")
	for i, v:=range b{
		fmt.Printf("b[%d]=%d\n",i,v)
	}
}

