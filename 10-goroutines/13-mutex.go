package main

import (
	"fmt"
	"sync"
)

// go run -race 13-mutex.go
var x int = 1

func main() {
	wg := new(sync.WaitGroup)
	m := new(sync.Mutex)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go UpdateX(i, wg, m)
	}
}

func UpdateX(val int, wg *sync.WaitGroup, m *sync.Mutex) {

	// data race situations
	//	- at least one concurrent write and n number of reads
	//	- n number of concurrent writes
	// 	- n number of concurrent writes and n number of concurrent reads
	// 	Note - Data race doesn't happen if there are only concurrent reads
	defer wg.Done()
	m.Lock()
	//the mutex is not locking a specific variable its locking a block of code
	defer m.Unlock()
	x = val
	fmt.Println(x)
}
