package main

import (
	"fmt"
	"os/user"
	"time"
)

func init() {
	fmt.Println("Init!")
}

func bazz() {
	fmt.Println("bazz!")
}

func main() {
	bazz()
	fmt.Println("Hello World", time.Now())
	fmt.Println(user.Current())
	var list []int
	list = make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		list = append(list, i)
		fmt.Println(list)
	}
	fmt.Println(list)
}
