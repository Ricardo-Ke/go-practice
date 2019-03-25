package main

import (
	"fmt"
)

// func sum(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		time.Sleep(1000 * time.Millisecond)
// 		sum += v
// 	}
// 	c <- sum
// }

// func fibonacci(n int, c chan int) {
// 	x, y := 0, 1
// 	for i := 0; i < n; i++ {
// 		c <- x
// 		x, y = y, x+y
// 		fmt.Printf("x = %d, y = %d\n", x, y)
// 	}
// 	close(c)
// }

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
			case 
		}
	}
}

func main() {
	// s := []int{1, 2, 3, 4, 5, 6}

	// c := make(chan int)
	// go sum(s[:len(s)/2], c)
	// go sum(s[len(s)/2:], c)
	// x, y := <-c, <-c

	// fmt.Println(x, y, x+y)

	// c := make(chan int, 10)
	// go fibonacci(cap(c), c)
	// for i := range c {
	// 	fmt.Println(i)
	// }


	
}
