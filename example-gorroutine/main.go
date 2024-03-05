package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	init := time.Now()

	var wg sync.WaitGroup
	var mu sync.Mutex

	count1 := 1
	count2 := 1

	for i := 1; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("com goroutine ", i)
			mu.Lock()
			count1++
			mu.Unlock()
		}(i)
		fmt.Println("sem goroutine ", i)
		mu.Lock()
		count2++
		mu.Unlock()
	}

	end := time.Now()
	wg.Wait()

	defer func() {
		fmt.Println("Contagem com goroutine: ", count1)
		fmt.Println("Contagem sem goroutine: ", count2)
	}()

	fmt.Println(end.Sub(init))
}
