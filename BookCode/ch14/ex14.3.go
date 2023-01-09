package main

import "fmt"

type Data struct {
	value int
	data [200]int
}

func ChangeData(arg Data){
	arg.value = 999
	arg.data[100] = 999
}

func ChangeData1(arg *Data){
	arg.value = 8888
	arg.data[100] = 8888
}


func main() {
	var data Data

	ChangeData(data)

	fmt.Printf("value = %d\n",data.value)
	fmt.Printf("data[100] = %d\n",data.data[100])

	ChangeData1(&data)
	fmt.Printf("value = %d\n",data.value)
	fmt.Printf("data[100] = %d\n",data.data[100])

}