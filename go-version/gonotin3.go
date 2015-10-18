/*
AUTHOR: ANDREW (GONUTS)
*/

package main

import (
	"os"
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func reader(r io.Reader) func() (uint64, error) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	return func() (uint64, error) {
		if !s.Scan() {
			if err := s.Err(); err != nil {
				return 0, err
			}
			return 0, io.EOF
		}
		return btou(s.Bytes())
	}
}

func btou(b []byte) (uint64, error) {
	var n uint64
	for _, v := range b {
		if v < '0' || v > '9' {
			return 0, strconv.ErrSyntax
		}
		if n >= ^uint64(0)/10+1 {
			return ^uint64(0), strconv.ErrRange
		}
		n *= 10
		n1 := n + uint64(v-'0')
		if n1 < n {
			return ^uint64(0), strconv.ErrRange
		}
		n = n1
	}
	return n, nil
}

func check2(err error) {
	if err != nil {
		die(err)
	}
}

func die(v ...interface{}) {
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(1)
}

func main() {
	f0, err := os.Open(os.Args[1])
	check2(err)
	defer f0.Close()

	f1, err := os.Open(os.Args[2])
	check2(err)
	defer f1.Close()

	m := make(map[uint64]struct{})
	r0 := reader(f1)
	for {
		n, err := r0()
		if err != nil {
			if err == io.EOF {
				break
			}
			die(err)
		}
		m[n] = struct{}{}
	}

	var b [20]byte
	r1 := reader(f0)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for {
		n, err := r1()
		if err != nil {
			if err == io.EOF {
				return
			}
			die(err)
		}
		if _, ok := m[n]; ok {
			continue
		}
		m[n] = struct{}{}
		w.Write(strconv.AppendUint(b[:0], n, 10))
		w.WriteByte('\n')
	}
}
