package go_utils

import (
	"fmt"
	"testing"
	"time"
)

func TestNewWaitGroup(t *testing.T) {
	wg := NewWaitGroup(10)
	for i := 0; i < 10; i++ {
		wg.Add()
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
			time.Sleep(time.Duration(2) * time.Second)
		}(i)
	}
	wg.Wait()
}
