package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	var n time.Duration
	fmt.Print("Введите время для работы программы (сек): ")
	fmt.Scan(&n)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*n)
    defer cancel()
	data := make(chan int)

	go func() {
		for {
			select{
			case <-ctx.Done():
                return
            default:
                var val int
                fmt.Scan(&val)
                data <- val
			} 
		}
	}()

    go func() {
        for {
            select{
            case <-data:
                return
            default:
                fmt.Println(<-data)
            }
        }
    }()
    <-ctx.Done()
    fmt.Println("Время вышло!")
}