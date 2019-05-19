package main

import "fmt"

func main() {
	//変数の中身とアドレスを確認する
	var number int = 100
	fmt.Println("content of variable:", number)
	fmt.Println("address of variable:", &number)

	//変数のアドレスを代入
	var pointer *int = &number
	fmt.Println("substitute number for pointer:", pointer)
	fmt.Println("content of pointer:", *pointer)
}
