package main

import (
	"fmt"
	"sync"
)

func SyncMap_Test1() {
	var m sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		tmp := i
		go func() {
			defer wg.Done()
			m.Store(tmp%10, tmp)
		}()
	}

	wg.Wait()
	m.Range(func(key, value any) bool {
		fmt.Println("key:", key, ", value:", value)
		return true // continue the iteration
	})

}
