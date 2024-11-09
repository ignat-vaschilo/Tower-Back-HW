package main

import (
    "fmt"
)

func setBit(num int64, i int, value int) int64 {
    if value == 1 {
        // Устанавливаем i-й бит в 1 с помощью операции побитового OR
        num |= (1 << i)
    } else if value == 0 {
        // Устанавливаем i-й бит в 0 с помощью операции побитового AND с инверсией
        num &= ^(1 << i)
    }
    return num
}

func main() {
    var num int64
    var i, value int

    fmt.Print("Введите число: ")
    fmt.Scan(&num)
    
    fmt.Print("Введите позицию бита для изменения (0-63): ")
    fmt.Scan(&i)

    fmt.Print("Введите значение бита (0 или 1): ")
    fmt.Scan(&value)

    if i < 0 || i > 63 || (value != 0 && value != 1) {
        fmt.Println("Неверные данные. Позиция бита должна быть от 0 до 63, значение - 0 или 1.")
        return
    }

    result := setBit(num, i, value)
    fmt.Printf("Результат: %d\n", result)
}