##  17장 숫자 맞추기 게임 만들기

rand()는 랜덤값 생성지원,
현대 컴퓨터로 완전한 랜덤 숫자를 구현하기 어려움 따라서 
Seed값을 현재 시각각으로 주어 유사 랜덤값을 산출한다.

```go
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
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
	rand.Seed(time.Now().UnixNano())
	
	r:=rand.Intn(100)//랜덤 값 생성
	cnt := 1
	for {
		fmt.Printf("type your number")
		n, err :=InputIntValue()
		if err != nil{
			fmt.Println("숫자만 입력")
		}else{
			if n>r{
				fmt.Println("n>r")
			}else if n<r{
				fmt.Println("n<r")
			}else {
				fmt.Println("정답 시도한 횟수 :",cnt, "입력한 숫자 : ", n)
				break
			}
		cnt++
		}
	}
}
```