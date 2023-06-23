package goleakalert

import (
	"fmt"
	"runtime"
	"time"
	"github.com/guptarohit/asciigraph"
)

// Detect Goroutine Leak
type LeakChecker struct {
	startNumberOfGoroutine 	int
	goroutineData 		 	[]float64
	done 					chan bool
}

func (l *LeakChecker) start() {
	// ticker := time.NewTicker(1 * time.Millisecond)
 	// defer ticker.Stop()

	for {
		select {
		case <- l.done:
			return
		case <- time.After(1 * time.Millisecond):
			fmt.Println("oi")
			l.goroutineData = append(l.goroutineData, float64(runtime.NumGoroutine()))
		}
	}
}

func (l *LeakChecker) stop() {
	fmt.Println("======= LeakChecker Start =======")
	l.done <- true

	// show goroutine graph
	graph := asciigraph.Plot(l.goroutineData)
	fmt.Println(graph)

	if l.startNumberOfGoroutine == runtime.NumGoroutine() {
		/*
			when there are no goroutine leak
		*/
		fmt.Println("No goroutine leaks")
	} else {
		/*
			when there are some goroutine leak
		*/
		panic("This code may cause goroutine leak")
	}

	fmt.Println("======= LeakChecker End =======")
}
