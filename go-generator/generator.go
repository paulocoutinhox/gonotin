package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func usage() {
	fmt.Printf("> usage: generator filename count max\n")
	os.Exit(1)
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	// version 1 : Paulo Coutinho
	// version 2 : Michael Jones (go-nuts)

	if len(os.Args) != 4 {
		usage()
	}
	filename := os.Args[1]
	count, e2 := strconv.Atoi(os.Args[2])
	max, e3 := strconv.Atoi(os.Args[3])
	if len(filename) < 1 || count < 0 || e2 != nil || max < 0 || e3 != nil {
		usage()
	}

	// generate a different sequence every time (if desired)
	rand.Seed(int64(time.Now().Nanosecond()))

	// create output file
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	buffer := bufio.NewWriter(file)
	for i := 0; i < count; i++ {
		fmt.Fprintln(buffer, random(0, max))
	}
	buffer.Flush()
	file.Close()

	fmt.Printf("> wrote %d numbers in (0 ≤ n < %d) to file %q\n", count, max, filename)
}