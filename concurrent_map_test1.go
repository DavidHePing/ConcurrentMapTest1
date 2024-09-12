package main

import (
	"fmt"

	cmap "github.com/orcaman/concurrent-map/v2"
)

func Concurrent_map_test1() {

	// Create a new map.
	m := cmap.New[string]()

	// Sets item within map, sets "bar" under key "foo"
	m.Set("foo", "bar")
	// m.Set("foo", 123) // compile error

	// Retrieve item from map.
	bar, ok := m.Get("foo")

	if ok {
		fmt.Println(bar)
	}

	// Removes item under key "foo"
	m.Remove("foo")

}
