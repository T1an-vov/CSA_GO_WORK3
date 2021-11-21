package main
import (
"fmt"
)
var myres = make(map[int]int, 20)
var mymutex = make(chan struct{},1)
func fa(n int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res*=i
	}
	myres[n]=res
	mymutex <- struct{}{}
}

func main() {
	for i := 1; i <= 20; i++ {
		go fa(i)
		<-mymutex
	}
	for i,v:=range myres{
		fmt.Printf("myres[%d]=%d\n",i,v)
	}
}
