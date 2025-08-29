package main

import "unsafe"

type Bocchi struct {
	x uint32
	y uint64
}

// step3.go
var (
	// assert Bocchi.x doesn't move
	_ [0]struct{} = [unsafe.Offsetof(Bocchi{}.x) - 0]struct{}{}
	// assert Bocchi.y comes next after Bocchi.x
	_ [0]struct{} = [unsafe.Offsetof(Bocchi{}.y) - (unsafe.Alignof(Bocchi{}.y))]struct{}{}
	// assert y is the last field in Bocchi
	_ [0]struct{} = [unsafe.Sizeof(Bocchi{}) - (unsafe.Alignof(Bocchi{}.y) + unsafe.Sizeof(Bocchi{}.y))]struct{}{}
)

func main() {}
