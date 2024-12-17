package gosyncmap

import (
	"sync"
	"testing"
)

func TestChannelMap_Goroutine(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(10)
	chanMap := NewChannelMap()
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				chanMap.Store("key", j)
			}
		}()
	}
	wg.Wait()
}
