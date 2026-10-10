package example3

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()         // Захоплення виключного доступу
	defer c.mu.Unlock() // Гарантоване звільнення після виходу
	c.value++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

//rename to main when running
func main3() {
	counter := SafeCounter{}
	var wg sync.WaitGroup
	const routinesCount = 1000
	wg.Add(routinesCount)
	for i := 0; i < routinesCount; i++ {
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}
	wg.Wait() // Очікування виконання всіх 1000 горутин
	fmt.Printf("Фінальне значення лічильника: %d\n", counter.Value()) // Завжди 1000
}