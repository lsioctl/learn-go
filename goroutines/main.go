package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)

	for i := 0; i < 10000; i++ {
		go func(i int) {
			time.Sleep(8 * time.Second)
			//fmt.Println(i)
			c <- i
		}(i)
	}

	for range c {
		fmt.Println(<-c)
	}

}
