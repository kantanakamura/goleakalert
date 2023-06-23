package goleakalert

import (
	"testing"
	"sync"
	"fmt"
	"runtime"
	"time"
)

func Test1(t *testing.T) {
	leakDetect := LeakDetect{runtime.NumGoroutine()}
	defer leakDetect.stop()

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

func Test2(t *testing.T) {
	leakDetect := LeakDetect{runtime.NumGoroutine()}
	defer leakDetect.stop()

	go func() {
		time.Sleep(10 * time.Second) 
	}()
	
	fmt.Println("hello world")
}