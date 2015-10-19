package main

/*
AUTHOR: MATT HARDEN (GONUTS)
DESCRIPTION: Using channel
*/

import (
	"bufio"
	"io"
	"log"
	"os"
)

func processFile(name string, work func(string), finis func()) {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		work(string(b)) // perform caller's work upon this string
	}
	finis()
}

func main() {
	// read the new "numbers" (strings) to test
	test := make(chan string, 1000000) // arbitrary value
	go processFile(os.Args[1], func(s string) { test <- s }, func() { close(test) })

	// read the known  "numbers" (strings) to test against
	reference := make(map[string]struct{})
	processFile(os.Args[2], func(s string) { reference[s] = struct{}{} }, func() {})

	// report unmatched numbers
	w := bufio.NewWriter(os.Stdout)
	for v := range test {
		_, found := reference[v]
		if !found {
			w.WriteString(v)
			w.WriteByte('\n')
			reference[v] = struct{}{} // uncomment for unique output
		}
	}
	w.Flush()
}
