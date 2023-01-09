package main

import (
	"fmt"
	"bufio"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error){

	var n int

	_, err := fmt.Scanln(&n)

	if err != nil{
		//input stream clear
		stdin.ReadString('\n')
	}
	return n, err
}

func main(){
	for {
		fmt.Printf("type your number")
		n, err :=InputIntValue()
		if err != nil{
			fmt.Println("숫자만 입력")
			fmt.Println("에러 :",err)
		}else{
			fmt.Println("입력하신 숫자는 ", n,"이다.")
		}
	}
}