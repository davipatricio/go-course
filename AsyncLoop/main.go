package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			fmt.Println(i)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 11000; i < 20000; i++ {
			fmt.Println(i)
		}
	}()

	// Wait for all goroutines to finish.
	wg.Wait()
	fmt.Println("Done counting!")

}
