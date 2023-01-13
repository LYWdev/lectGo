package main

import (
	"fmt"
	"os"
)

func main (){
	h := []byte("Hello")
	f, err := os.Create("test.txt")
	
	if err != nil{
		fmt.Println("Failed to create")		
		return
	}
	defer fmt.Println("Must Call")
	defer f.Close()
	defer fmt.Println("File Close")

	fmt.Println("Write Hello World")
	fmt.Println(f, string(h))
}

