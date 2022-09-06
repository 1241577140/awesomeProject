package main

import (
	"fmt"
	"time"
)

type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	mu := &Mutex{ch: make(chan struct{}, 1)}
	mu.ch <- struct{}{}
	return mu
}

func (m *Mutex) Lock() {
	<-m.ch
}

func (m *Mutex) TryLock() bool {
	select {
	case <-m.ch:
		return true
	default:
		return false
	}
}

func (m *Mutex) TryLockWithTimeout(t time.Duration) bool {
	c := time.NewTimer(t)
	select {
	case <-c.C:
		return false
	case <-m.ch:
		return true
	}
}

func (m *Mutex) Unlock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("如果没写进去，代表重复解锁")
	}
}
func main() {
	mutex := NewMutex()
	mutex.Lock()
	tryLock := mutex.TryLock()
	fmt.Println(tryLock)
	tlt := mutex.TryLockWithTimeout(time.Second)
	fmt.Println(tlt)
	mutex.Unlock()
}
