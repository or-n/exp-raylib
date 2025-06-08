package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	max    = 100000
	p      = make([]int, max)
	q      = make([]int, max)
	powers = make([]int, max)
	m      = 998244353
)

func line() string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func Int(i string) int {
	x, _ := strconv.Atoi(i)
	return x
}

func choose(p1, q1, p2, q2 int) (int, int) {
	if p1 != q2 {
		if p1 > q2 {
			return p1, q1
		}
		return p2, q2
	}
	if q1 > p2 {
		return p1, q1
	}
	return p2, q2
}

func main() {
	x := 1
	for i := range max {
		powers[i] = x
		x = (x * 2) % m
	}
	tests := Int(line())
	for range tests {
		n := Int(line())
		for i, part := range strings.Fields(line()) {
			p[i] = Int(part)
		}
		for i, part := range strings.Fields(line()) {
			q[i] = Int(part)
		}
		max_p, max_q := 0, 0
		for i := range n {
			if p[i] > p[max_p] {
				max_p = i
			}
			if q[i] > q[max_q] {
				max_q = i
			}
			a, b := choose(p[max_p], q[i-max_p], p[i-max_q], q[max_q])
			r := (powers[a] + powers[b]) % m
			fmt.Print(r, " ")
		}
		fmt.Println()
	}
}
