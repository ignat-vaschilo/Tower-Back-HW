package main

import (
	"context"
	"fmt"
	"time"
)

func cancelWithCtx(ctx context.Context) {

	for {
		select {

		default:
			fmt.Println("работаю")

		case <-ctx.Done():
			fmt.Println("завершаю работу")
			return
		}
	}

}

func cancelWithChan(stop chan bool) {

	for {
		select {

		default:
			fmt.Println("работаю")

		case <-stop:
			fmt.Println("завершаю работу")
			return
		}
	}

}

func main() {
	stop := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())

	go cancelWithChan(stop)
	go cancelWithCtx(ctx)

	time.Sleep(2 * time.Millisecond)
	// с помощью отправки стоп сигнала в канал
	stop <- true
	// через cancel
	cancel()
}