package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func GenNums(size int, genRange int) []int {
	if size <= 0 {
		panic("size must be positive!")
	}
	randSrc := rand.NewSource(time.Now().Unix())
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = int(randSrc.Int63() % int64(genRange))
	}
	return data
}

func CreateRoutines() {
	var wg sync.WaitGroup
	fromGen := make(chan int)
	toMain := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		data := GenNums(10, 100)
		for _, v := range data {
			fromGen <- v
		}
		close(fromGen)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range fromGen {
			toMain <- v * v
		}
		close(toMain)
	}()
	for v := range toMain {
		fmt.Print(v, " ")
	}
	wg.Wait()
	fmt.Println()

}
