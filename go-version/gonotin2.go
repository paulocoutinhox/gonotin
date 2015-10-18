/*
AUTHOR: JUSTIN (GONUTS)
*/

package main

import (
	"os"
	"log"
	"bufio"
	"io"
)

func readMap(filename string) map[string]struct{} {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	m := make(map[string]struct{})
	empty := struct{}{}
	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		m[string(b)] = empty
	}
	return m
}

func main() {
	// read files
	mapA := readMap(os.Args[1])
	mapB := readMap(os.Args[2])

	w := bufio.NewWriter(os.Stdout)

	// process and show data from A that not exists in B
	for item := range mapA {
		if _, ok := mapB[item]; !ok {
			w.WriteString(item)
			w.WriteByte('\n')
		}
	}

	w.Flush()
}
