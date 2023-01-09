package main

import (
	"fmt"
	"unsafe"
)

//구조제 정의
type User1 struct {
	A int8	    // 4바이트
	B int	    // 4바이트
	C int8	    // 4바이트
	D int	    // 4바이트
	E int8	    // 4바이트
}

type User2 struct {
	A int8	    // 1바이트
	C int8	    // 1바이트
	E int8	    // 4바이트
	B int	    // 8바이트
	D int	    // 8바이트
}

func main() {
	user := User1{1,2,3,4,5} 
	user2 := User2{6,7,8,9,10} 
    fmt.Println(unsafe.Sizeof(user))
    fmt.Println(unsafe.Sizeof(user2))
}