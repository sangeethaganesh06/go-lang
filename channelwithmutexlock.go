package main

import (
	"fmt"
	"sync"
)

const no = 10

func main() {
	num := 0
	mutex := sync.Mutex{}
	done := make(chan struct{})

	go func() {
		for i := 0; i < no; i++ {
			mutex.Lock()
			num++
			mutex.Unlock()
		}
		done <- struct{}{}
	}() //for call the go routine
	go func() {
		for i := 0; i < no; i++ {
			mutex.Lock()
			num--
			mutex.Unlock()
		}
		done <- struct{}{}
	}()
	<-done //to read two  value coz its using single go routine
	<-done
	fmt.Println(num)
}
