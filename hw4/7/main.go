package main

import (
	"fmt"
	"sync"
)

type ConccurentMap struct {
	nums map[int]int
	sync.RWMutex
}

func (c *ConccurentMap) WriteToMap(data int, i int) {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()
	c.nums[i] = data
}

func main() {
	var wg sync.WaitGroup

	numbers := &ConccurentMap{
		nums: map[int]int{},
	}

	for i := 0; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			numbers.WriteToMap(i, i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Print(numbers.nums)
}