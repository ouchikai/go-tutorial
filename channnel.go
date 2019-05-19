package main

import "fmt"

func gochannel(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go gochannel(s, c)
	x := <-c
	fmt.Println(x)
}
