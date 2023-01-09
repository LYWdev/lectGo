package main

import (
	"fmt"
	"unsafe"
)

//구조제 정의
type User struct {
	Age int32	    // 4바이트
	Score float64   // 8바이트
}

func main() {
	user := User{23,77.2} 
    fmt.Println(unsafe.Sizeof(user))
    // int 32비트 -> 8비트 : 1 바이트 즉 4바이트
    // (64+32)/8 -> 12바이트가 나와야함
    // 그런데 실제 구동하면 16바이트가 나옴.
    // 메모리 정렬 때문임.
    // 메모리 정렬이란?


}