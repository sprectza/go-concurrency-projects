package main

import (
	"fmt"
	"sync"
)

func printsomething(s string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrements wait group by 1
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
		"nu",
	}

	wg.Add(len(words)) // 10 since 10 words are there

	for i, x := range words {
		// waitgroups should always be passed around by address, not copy
		go printsomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait()

	wg.Add(1)
	printsomething("This is the second thing to be printed", &wg)
}
