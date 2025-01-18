package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	cmap "github.com/orcaman/concurrent-map/v2"
)

func Concurrent_map_test2() {

	m := cmap.New[int]()

	var wg sync.WaitGroup
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		tmp := i
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			m.Set(strconv.Itoa(tmp%10), tmp)
		}()
	}

	wg.Wait()

	for key, value := range m.Items() {
		fmt.Println("key:", key, ", value:", value)
	}
}
