package main

import (
	"fmt"
	"sync"
)
var wg = sync.WaitGroup{}

func A() {
	fmt.Print("A")
	wg.Done()
}
func B() {
	fmt.Print("B")
	wg.Done()
}
func C() {
	fmt.Print("C")
	wg.Done()
}
func main() {
	for i := 0; i < 10; i++ {
		go A()
		wg.Add(1)
		wg.Wait()
		go B()
		wg.Add(1)
		wg.Wait()
		go C()
		wg.Add(1)
		wg.Wait()
		fmt.Print("\n")
	}
}

