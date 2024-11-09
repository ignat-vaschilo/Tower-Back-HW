package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func worker(i int, data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range data {
		fmt.Printf("Воркер %d выводит значение %d\n", i, val)
	}
}

func main() {
	var cntWorkers int
	fmt.Println("Введите количество воркеров: ")
	fmt.Scan(&cntWorkers)

	data := make(chan int)
	done := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(cntWorkers)
	for i := 0; i < cntWorkers; i++ {
		go worker(i+1, data, &wg)
	}

	go func() {
		for {
			select {
			case <-done:
				close(data)
				return
			default:
				val := rand.Intn(100)
				data <- val
			}
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Println("\nЗаканчиваем воркеры")
		close(done)
	}()

	wg.Wait()
	fmt.Println("\nВсе воркеры закрылись")
}
