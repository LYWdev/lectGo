package interfaces
import (
	"fmt"
	"io"
	"os"
)

//Copy 함수는 버퍼를 사용해 in ->  out의 첫번째 dir로 데이터 복사
//stdout에도 data를 write
func Copy(in io.ReadSeeker, out io.Writer)error{
	//매개변수  out와 표준 출력(Stdout)에도 데이터를 쓴다.
	w:=io.Multwriter(out,os.Stdout)

	if _, err :=io.Copy(w,int);err!=nill{
		return err
	}
	in.Seek(0,0)
	Buf :=make([]byte, 64)
	if_, err :=io.CopyBuffer(w,in,buf);err!=nill{
		return err
	}
	fmt Println()
	return nil
}