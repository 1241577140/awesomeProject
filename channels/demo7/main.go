package main

func main() {
	var ch chan struct{}
	ch <- struct{}{}
}
