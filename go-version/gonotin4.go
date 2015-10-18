package main

/*
AUTHOR: MICHAEL JONES (GONUTS)
*/

import (
	"bufio"
	"io"
	"log"
	"os"
	"sync"
)

func main() {
	mode8string()
}

func mode8string() {
	var wg sync.WaitGroup
	wg.Add(2)

	// read the new "numbers" (strings) to test
	test := make([]string, 0, 1000000) // arbitrary value to avoid small reallocations
	go func() {
		f, err := os.Open(os.Args[1])
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

			test = append(test, string(b))
		}
		wg.Done()
	}()

	// read the known  "numbers" (strings) to test against
	reference := make(map[string]struct{})
	go func() {
		f, err := os.Open(os.Args[2])
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

			reference[string(b)] = struct{}{}
		}
		wg.Done()
	}()

	// wait until both files have been read
	wg.Wait()

	// report unmatched numbers
	w := bufio.NewWriter(os.Stdout)
	for _, v := range test {
		_, exists := reference[v]
		if !exists {
			w.WriteString(v)
			w.WriteByte('\n')
		}
	}
	w.Flush()
}