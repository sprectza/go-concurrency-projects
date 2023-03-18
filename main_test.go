package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

// Keep the T capital in test
func Test_printsomething(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup

	wg.Add(1)

	go printsomething("eta", &wg)

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "eta") {
		t.Errorf("Exepected to find eta, but not found")
	}
}
