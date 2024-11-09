package main

import (
	"fmt"
	"sync"
)

const (
	n = 10
)

func main() {
	input := make(chan int)
	output := make(chan int)
	arr := [n]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var wg sync.WaitGroup
	
	wg.Add(3)
	go func() {
		for i := 0; i < n; i++ {
			input <- arr[i]
		}
		close(input)
		wg.Done()
	}()

	go func() {
		for i := range input {
			output <- i*i
		}
		close(output)
		wg.Done()
	}()
	
	go func() {
		for i := range output {
			fmt.Println(i)
		}
		wg.Done()
	}()
	
	wg.Wait()
}