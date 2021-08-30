package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()

	for i := 0; i < 5; i++ {
		x, ok := <-ch
		fmt.Println(ok, x)
	}
	/*
		true 0
		true 1
		true 2
		true 3
		true 4
	*/
	fmt.Println("======================")

	ch2 := make(chan int)
	go func() {
		defer close(ch2)
		for i := 0; i < 5; i++ {
			ch2 <- i
		}
	}()

	for i := 0; i < 6; i++ {
		x, ok := <-ch2
		fmt.Println(ok, x)
	}
	/*
		true 0
		true 1
		true 2
		true 3
		true 4
		false 0
	*/
	fmt.Println("======================")

	/*ch3 := make(chan int) // commented as it produces deadlocks
	go func() {
		//defer close(ch3)
		for i := 0; i < 5; i++ {
			ch3 <- i
		}
	}()

	for i := 0; i < 6; i++ {
		x, ok := <-ch3
		fmt.Println(ok, x)
	}
	/*
		true 0
		true 1
		true 2
		true 3
		true 4
		fatal error: all goroutines are asleep - deadlock!
	*/
	fmt.Println("======================")

	ch4 := make(chan int)
	go func() {
		//defer close(ch3)
		for i := 0; i < 5; i++ {
			ch4 <- i
		}
	}()

	for i := 0; i < 4; i++ {
		x, ok := <-ch4
		fmt.Println(ok, x)
	}
	/*
		true 0
		true 1
		true 2
		true 3
	*/
}
