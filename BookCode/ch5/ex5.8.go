package main

import (
	"bufio"		        // IO를 담당하는 패키지	
	"fmt"	
	"os"   		        // IO를 담당하는 패키지
)

func main() {
	stdin := bufio.NewReader(os.Stdin)//표준 입력을 읽는 객체

	var a int
	var b int

	n,err:=fmt.Scanln(&a,&b)
	if err != nil{						//에러발생시
		fmt.Println(err)				//에러 출력
		stdin.ReadString('\n') 			//Scanln 은 입력 끝애 개행이 있음.개행이 나올때 까지 읽어라는 의미임 즉, 한번 쭉 읽으면 버퍼가 지워짐.
	}else{
		fmt.Println(n,a,b)
	}

	n,err =fmt.Scanln(&a,&b)
	if err != nil{						//에러발생시
		fmt.Println(err)				//에러 출력
	}else{
		fmt.Println(n,a,b)
	}

}
