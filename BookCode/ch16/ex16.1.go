package main

// 모듈은 하나 이상의 패키지의 묶음을 의미함
// 패키지란 코드를 묶는 단위
import (
	"fmt"
	//경로가 있는 패키지 
	"math/rand"
	"text/template"
	htemplate "html/template"
	//사용하지 않는 패키지를 임포트하는 경우
	_ "html/template"
)


func main(){
	//경로가 있는 패키지는 마지막 폴더명 rand만 사용한다.
	fmt.Println(rand.Int())
	//경로가 겹칠 때, 별칭으로 구분
	// template.Println(rand.Int())
}