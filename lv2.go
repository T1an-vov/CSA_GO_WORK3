package main
import (
	"fmt"
	"runtime"
)
func putNum(intChan chan int) {
	for i := 1; i <= 50000; i++ {
		intChan<- i
	}
	close(intChan)
}
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num % i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan<- num
		}
	}
	exitChan<- true
}
func main() {
	runtime.GOMAXPROCS(10)
	intChan := make(chan int , 1000)
	primeChan := make(chan int, 50000)
	exitChan := make(chan bool, 8)
	go putNum(intChan)
	for i := 0; i < 10; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}
	go func(){
		for i := 0; i < 10; i++ {
			<-exitChan
		}
		close(primeChan)
	}()
	for {
		res	, ok := <-primeChan
		if !ok{
			break
		}
		fmt.Println(res)
	}
}

