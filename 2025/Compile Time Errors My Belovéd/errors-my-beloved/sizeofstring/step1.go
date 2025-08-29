//go:build ignore

// step1.go
package main

import "unsafe"

type Bocchi struct {
	x uint32
	y uint64
}

var _ [0]struct{} = [unsafe.Sizeof(Bocchi{}) - 0]struct{}{}

func main() {}
