package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(5)
	sm := 0
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			sm += arr[i] * arr[i]
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println(sm)
}
