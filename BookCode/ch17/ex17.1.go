package main

import (
	"fmt"
	"math/rand"
	"time"
)

var stdin = bufio.NewReader(os.stdin)

func InputIntValue()(int,error){
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil{
		stdin.ReadString('\n')
	}
	return n, err
}

func main(){
	for {
		fmt.Printf("숫자 값을 입력하세요")
		n,err :=InputIntValue()
		if err != nil {
			fmt.Println("숫자값을 입력하세요")
		}
	}
}