package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)

		go func (i int) {
			defer wg.Done()
			fmt.Println(i)
			time.Sleep(1)
		}(i)
	}
	wg.Wait()
	fmt.Println("complete!")
}