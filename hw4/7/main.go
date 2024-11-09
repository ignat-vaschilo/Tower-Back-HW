package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	hash := make(map[string]int)

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			hash[fmt.Sprintf("key%d", i)] = i
		}(i)
	}
	wg.Wait()
	fmt.Println(hash)
}
