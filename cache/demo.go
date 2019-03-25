package main

import (
	"fmt"
	"sync"
)

var count int = 0

func main() {
	var wg sync.WaitGroup
	var t sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(idx int) {
			t.Lock()
			defer t.Unlock()

			fmt.Println("index is : ", idx, "count now is: ", count)
			count++
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println(count)
}
