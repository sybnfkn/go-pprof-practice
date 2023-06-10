package demo

import (
	"errors"
	"fmt"
	"time"
)

func TestGoroutine() {
	for {
		err := handle()
		if err != nil {
			fmt.Println(err.Error())
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func handle() error {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func(i int) {
			if i == 0 {
				time.Sleep(1500 * time.Millisecond)
			} else {
				time.Sleep(500 * time.Millisecond)
			}
			ch <- i
		}(i)
	}
	time := time.After(1 * time.Second)
	count := 0
	for count < 10 {
		select {
		case <-ch:
			count += 1
		case <-time:
			return errors.New("timeout")
		}
	}
	return nil
}
