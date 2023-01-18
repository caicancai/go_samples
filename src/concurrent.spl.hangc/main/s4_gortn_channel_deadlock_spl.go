package main

import "fmt"

func main() {
	deadLockSpl()

}

// 死锁演示
func deadLockSpl() {
	ch := make(chan int)
	ch <- 1
	a := <-ch
	fmt.Println(a)

}

// deadLockSpl 避免死锁
func deLockSpl() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	a := <-ch
	fmt.Println(a)
	/**
	1
	*/
}
