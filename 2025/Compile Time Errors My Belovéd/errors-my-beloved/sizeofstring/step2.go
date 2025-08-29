//go:build ignore

// step2.go
package main

import "unsafe"

type Bocchi struct {
	x uint32
	y uint64
}

var _ [0]struct{} = [unsafe.Sizeof(Bocchi{}) - 16]struct{}{}

func main() {}
