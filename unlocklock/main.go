package main

import (
	"sync"
	"time"
)

var l sync.Mutex

func main() {
	l.Lock()
	go func() {
		time.Sleep(time.Second)
		l.Unlock()
	}()
	l.Lock()
}
