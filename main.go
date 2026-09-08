package main

import "sync"

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			wg.Add(1)
		}()
	}
	wg.Wait()
}
