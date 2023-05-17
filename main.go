package main

import (
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	log.Print("loaded WASM module")
	before := time.Now()
	n := 0
	for i := 0; i < 4; i++ {
		wg.Add(1)
		from := i * 40_000
		to := (i + 1) * 40_000
		go func() {
			log.Printf("Thread %d:%d started", from, to)
			n += thread(from, to)
			wg.Done()
		}()
	}

	wg.Wait()
	log.Printf("Total primes: %d", n)
	after := time.Now()
	log.Print("Total time: ", after.Sub(before))

	n = thread(0, 160_000)
	log.Printf("Total primes: %d", n)
}

func thread(from, to int) int {
	before := time.Now()
	n := prime(from, to)
	after := time.Now()
	log.Printf("Thread %d:%d -> %d primes in %v", from, to, n, after.Sub(before))
	return n
}

func prime(from, to int) int {
	count := 0
	for i := from; i < to; i++ {
		prime := false
		for j := 2; j < i; j++ {
			if i%j == 0 {
				prime = true
				break
			}
		}
		if prime {
			count++
		}
	}
	return count
}
