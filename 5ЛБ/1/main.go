package main

import (
	"fmt"
	"time"
)

func count(ch <-chan int) {
	for num := range ch {
		fmt.Printf("Квадрат числа %d: %d\n", num, num*num)
	}
}

func main() {
	ch := make(chan int)

	// Запуск функции count в отдельной горутине
	go count(ch)

	// Отправка чисел в канал
	for i := 1; i <= 5; i++ {
		ch <- i
	}

	// Закрытие канала
	close(ch)

	// Пауза для завершения работы горутины
	time.Sleep(time.Second)
}
