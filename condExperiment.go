package main

import (
	"fmt"
	"sync"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})

	i := 0

	gg := func() {
		fmt.Println("in2")
	}

	ff := func() {
		fmt.Println("in")
		c.L.Lock()
		fmt.Println(i)
		c.L.Unlock()
		c.Signal()
	}

	for i < 10 {
		c.L.Lock()
		if i == 5 {
			c.Wait()
		}
		i++
		go gg() // calls every time
		go ff() // but cannot execute the variable related operations without the wait being called
		c.L.Unlock()
	}
}

/*
sample output:
in
in
5
in2
in2
in
in
in2
in2
in

*/
