package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	arr := [5]int{2, 4, 6, 8, 10}
	for i := 0; i < 5; i ++ {
		go func(i int) {
			fmt.Printf("%d power of 2: %d\n", arr[i], arr[i]*arr[i])
			wg.Done()
		}(i)
	}
	wg.Wait()

}