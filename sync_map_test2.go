package main

import (
	"fmt"
	"sync"
)

func SyncMap_Test2() {
	var m sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		tmp := i
		go func() {
			defer wg.Done()
			if tmp%2 == 0 {
				m.Store(tmp%10, tmp)
			} else {
				m.Store(tmp%10, string(rune(tmp)))
			}
		}()
	}

	wg.Wait()
	m.Range(func(key, value any) bool {
		fmt.Println("key:", key, ", value:", value)
		return true // continue the iteration
	})

}
