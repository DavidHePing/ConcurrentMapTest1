package main

import (
	"fmt"
	"strconv"
	"sync"

	cmap "github.com/orcaman/concurrent-map/v2"
)

func Concurrent_map_test2() {

	m := cmap.New[int]()

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		tmp := i
		go func() {
			defer wg.Done()
			m.Set(strconv.Itoa(tmp%10), tmp)
		}()
	}

	wg.Wait()

	for key, value := range m.Items() {
		fmt.Println("key:", key, ", value:", value)
	}
}
