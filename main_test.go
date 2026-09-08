package main

import (
	"sync"
	"testing"
)

func TestMainFunc(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		t.Run("xxx", func(t *testing.T) {
			t.Parallel()
			func() {
				defer wg.Done()
			}()
		})
	}
	wg.Wait()

}
