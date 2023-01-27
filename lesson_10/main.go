package main

import (
	"fmt"
	"sync"
	//"sync/atomic"
	//"time"
)

// type Counter struct {
// 	mu sync.Mutex
// 	c  map[string]int
// }

// type Counter struct {
// 	mu sync.RWMutex
// 	c  map[string]int
// }

// func (c *Counter) Inc(key string) {
// 	c.mu.Lock()
// 	c.c[key]++
// 	c.mu.Unlock()
// }

// func (c *Counter) Value(key string) int {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.c[key]
// }

// func (c *Counter) CountMe() map[string]int{
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.c
// }

// func (c *Counter) CountMeAgain() map[string]int{
// 	c.mu.RLock()
// 	defer c.mu.RUnlock()
// 	return c.c
// }

func main() {
	// key := "test"
	// c := Counter{c: make(map[string]int)}
	// for i := 0; i < 1000; i++ {
	// 	go c.Inc(key)
	// }
	// time.Sleep(3 * time.Second)
	// fmt.Println(c.Value(key))
	var wg sync.WaitGroup
	var counter uint64
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		//k := i
		go func() {
			defer wg.Done()

			// fmt.Printf("%d goroutine working..\n", k)
			// time.Sleep(500 * time.Millisecond)

			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock() 
				// atomic.AddUint64(&counter, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Printf("all done. counter = %d\n", counter)
}
