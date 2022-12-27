package main

import "fmt"

func main (){
	// int 와 flaot64 자료형 변수 2개 선언
	a:=3
	var b float64 = 3.5

	var c int = int(b) //float 64에서 int 로 전홚
	d:=float64(a*c) // int 에서 float 64로 전환 int와 int 를 곱한 값을 float64로 전환하여 d에 저장
	
	var e int64 = 7 
	f := int64(d)*e //float64에서 int 64로 전환 

	var g int = int(b*3)//float64 에서 int로 변환
	var h int = int(b)*3//float64 에서 int로 변환

	fmt.Println(a, b,"c :%d "d,e,f,g,h,f,c)


}
