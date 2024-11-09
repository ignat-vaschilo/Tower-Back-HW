package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("by time")
	}()

	done := make(chan bool)
	go func() {
		fmt.Println("balancing")
		done <- true //отправляем
	}()
	<-done //ожидаем откуда то данных

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("context is done")
				return
			default:
				fmt.Println("context is aight")
			}
		}
	}()
	cancel()

	stop := make(chan struct{})
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("channel is closed so is the 4th")
				return
			default:
				fmt.Println("4th works")
			}
		}
	}()
	close(stop)

	timectx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	go func() {
		for {
			select {
			case <-timectx.Done():
				fmt.Println("time is up for 5th")
				return
			default:
				continue
			}
		}
	}()

	time.Sleep(2 * time.Second) // даем горутинкам запуститься
	fmt.Println("Main: All goroutines stopped")
}
