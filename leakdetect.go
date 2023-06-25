package goleakalert

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// get a goroutine id
func goid() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	// stateField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[1]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	// fmt.Println(stateField)
	return id
}

// Detect Goroutine Leak
type LeakDetect struct {
	startNumberOfGoroutine int
}

func (l *LeakDetect) stop() {
	fmt.Println("======= LeakDetect Start =======")

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

	fmt.Println("=======  LeakDetect End  =======")
}
