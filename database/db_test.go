package database

import (
	"fmt"
	"sync"
	"testing"
)

func TestName(t *testing.T) {
	count := 543
	groupSize := 100
	var wg sync.WaitGroup
	numGroups := (count-1)/groupSize + 1
	wg.Add(numGroups)

	for i := 0; i < count; i += groupSize {
		start := i
		end := i + groupSize - 1
		if end > count {
			end = count - 1
		}
		go func(start, end int) {
			defer wg.Done()
			for j := start; j <= end; j++ {
				fmt.Println(j)
			}
		}(start, end)
	}

	wg.Wait()
}
