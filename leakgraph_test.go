package goleakalert

import (
	"testing"
	"sync"
	"fmt"
	"runtime"
	"time"
)

func Test3(t *testing.T) {
	// leakChecker := LeakChecker{runtime.NumGoroutine(), []float64{1, 2, 4, 6, 10, 7, 4, 2, 1}, make(chan bool)}
	leakChecker := LeakChecker{runtime.NumGoroutine(), []float64{}, make(chan bool)}
	go leakChecker.start(100 * time.Microsecond)
	defer leakChecker.stop()

	fmt.Println("main", goid())
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i, goid())
		}()
	}
	wg.Wait()
}

