package gosyncmap

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRWMutexMap_Goroutine(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(10)
	rwMap := NewRWMutexMap()
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				rwMap.Store("key", j)
			}
		}()
	}
	wg.Wait()
}

func TestRWMutexMap_Store(t *testing.T) {
	rwMap := NewRWMutexMap()

	rwMap.Store("key1", "value1")
	value, ok := rwMap.Load("key1")
	assert.Equal(t, value, "value1")
	assert.True(t, ok)
}
