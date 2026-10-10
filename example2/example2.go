package example2

import (
	"fmt"
	"time"
)

func main2() {
	dataChan := make(chan string)

	// Імітація тривалої фонової роботи
	go func() {
		time.Sleep(2 * time.Second)
		dataChan <- "результат обчислень"
	}()

	// Очікування результату з жорстким таймаутом у 1 секунду
	select {
	case res := <-dataChan:
		fmt.Println("Отримано дані:", res)

	case <-time.After(1 * time.Second):
		// time.After повертає канал, у який приходить сигнал після заданого часу
		fmt.Println("Помилка: перевищено ліміт часу (Timeout)!")

	default:
		fmt.Println("Other case: Виконується інша робота, поки очікуємо результат...")
		// Спрацювало б одразу, якби не було потреби чекати
		// (використовується, коли блокування небажане)
	}
}