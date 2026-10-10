package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func printMemUsage(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// Переводимо байти в мегабайти
	fmt.Printf("[%s] Використано пам'яті (Heap + Stack): %.2f MB\n", 
		label, float64(m.Sys)/1024/1024)
}

func main() {
	// Кількість одночасних задач
	const totalGoroutines = 100_000
	fmt.Printf("Кількість логічних процесорів (P у GMP): %d\n", runtime.GOMAXPROCS(0))
	printMemUsage("До створення горутин")
	var wg sync.WaitGroup
	wg.Add(totalGoroutines)
	startTime := time.Now()
	for i := 0; i < totalGoroutines; i++ {
		// Запуск легкої горутини ключовим словом "go"
		go func(id int) {
			defer wg.Done()
			// Імітація легкої роботи / очікування події
			time.Sleep(100 * time.Millisecond)
		}(i)
	}
	// Фіксуємо пам'ять, коли всі 100 000 горутин активні
	time.Sleep(10 * time.Millisecond)
	printMemUsage("100 000 горутин активні")
	// Очікуємо завершення всіх завдань
	wg.Wait()
	fmt.Printf("Усі горутини успішно виконані за: %v\n", time.Since(startTime))
	// Звільнення ресурсів
	runtime.GC()
	printMemUsage("Після завершення та GC")
}