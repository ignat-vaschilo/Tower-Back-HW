package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int64
	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	fmt.Print("Введите i-ый бит: ")
	var i int
	fmt.Scan(&i)

	fmt.Print("Введите значение бита (0 или 1): ")
	var value string
	fmt.Scan(&value)

	binN := strconv.FormatInt(n, 2)
	if i < 0 || i >= len(binN) || (value != "0" && value != "1") {
		fmt.Println("Неверные данные")
		return
	}

	answer := binN[:i] + value + binN[i+1:]
	result, ok := strconv.ParseInt(answer, 2, 64)
	if ok == nil {
		fmt.Println("Результат:", result)
	}
}
