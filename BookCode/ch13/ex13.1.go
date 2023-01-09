package main

import "fmt"

//구조제 정의
type House struct {
	Address string
	Size int
	Price float64
	Type string
}

func main() {
	var house House
	house.Address = "사랑시 행복구 "
	house.Size =28 
	house.Price =9.8
	house.Type = "apt"

	fmt.Println("주소",house.Address)
	fmt.Println("크기: %d평",house.Size)
	fmt.Printf("가격: %.2f억원\n",house.Price)
	fmt.Println("주소",house.Type)
}