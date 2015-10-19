package main

/*
AUTHOR: MATT HARDEN (GONUTS)
DESCRIPTION: Using channel
*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type number uint64

func (n number) String() string {
	if n == 0 {
		return "0"
	}
	buf := make([]byte, 20)
	i := len(buf) - 1
	for ; n > 0; i-- {
		var d byte
		n, d = n/10, byte(n%10)
		buf[i] = d + '0'
	}
	return string(buf[i:])
}

type ErrNumberTooLong string

func (e ErrNumberTooLong) Error() string {
	return fmt.Sprintf("Number too long: %q", string(e))
}

type ErrNumberSyntax string

func (e ErrNumberSyntax) Error() string { return fmt.Sprintf("Invalid number: %q", string(e)) }

const maxnumber = ^number(0)

func ParseNumber(s string) (number, error) {
	var n number
	for _, d := range s {
		switch {
		case '0' <= d && d <= '9':
			if n > (maxnumber-9)/10 {
				return 0, ErrNumberTooLong(s)
			}
			n = 10*n + number(d-'0')
		case d == '(' || d == ')' || d == '+' || d == '-':
			// ignore
		default:
			return 0, ErrNumberSyntax(s)
		}
	}
	return n, nil
}

type numbers map[number]struct{}

func NewNumbers() numbers { return make(numbers, 100000) }

func (ns numbers) Add(n number) { ns[n] = struct{}{} }

func (ns numbers) Has(n number) bool {
	_, ok := ns[n]
	return ok
}

func processFile(name string, c chan<- number) {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		if n, err := ParseNumber(s.Text()); err != nil {
			log.Fatal(err)
		} else {
			c <- n
		}
	}
	if err := s.Err(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// read the new numbers to test
	test := make(chan number, 1000000) // arbitrary value
	go func() {
		defer close(test)
		processFile(os.Args[1], test)
	}()

	// read the known numbers to test against
	known := make(chan number, 1000000)
	go func() {
		defer close(known)
		processFile(os.Args[2], known)
	}()

	reference := NewNumbers()
	for v := range known {
		reference.Add(v)
	}

	// report unmatched numbers
	w := bufio.NewWriter(os.Stdout)
	for v := range test {
		if !reference.Has(v) {
			w.WriteString(v.String())
			w.WriteByte('\n')
			reference.Add(v)
		}
	}
	w.Flush()
}
