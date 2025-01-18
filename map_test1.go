package main

import (
	"fmt"
	"sync"
)

func Map_Test1_panic() {
	m := make(map[int]int)
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		tmp := i
		go func() {
			defer wg.Done()
			m[tmp%10] = tmp
		}()
	}

	wg.Wait()
	fmt.Println(m)

}
