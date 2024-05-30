package main

import (
	"context"
	"fmt"
	"sync"
)

// 利用筛法并行求素数

func selectPrime(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	prime, ok := <-c
	if !ok {
		return
	}
	fmt.Println(prime)
	newChan := make(chan int)
	newWg := sync.WaitGroup{}
	newWg.Add(1)
	go selectPrime(newChan, &newWg)
	for n := range c {
		if n%prime != 0 {
			newChan <- n
		}
	}
	close(newChan)
	newWg.Wait()
}

func printPrime(ctx context.Context, n int) {
	// 筛法求素数
	wg := sync.WaitGroup{}
	wg.Add(1)
	c := make(chan int)
	go selectPrime(c, &wg)
	for i := 2; i <= n; i++ {
		c <- i
	}
	close(c)
	wg.Wait()
}
