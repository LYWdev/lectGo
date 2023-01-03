package main

import "fmt"

func main() {
	var t[5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}
	//range는 배열을 순환하면서 각각 인덱스값과 데이터 값(i,v)을 반환한다.
	// 인덱스 값이 필요없다면 아래 처럼하면된다.
	// for _, v:=range t{
	for i, v:=range t{
		fmt.Println(i,v)

	}
}
