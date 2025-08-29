package main

type Guitarist interface {
	Strum(chord Chord)
}

type Bocchi struct {
	Guitar string
}

func (b *Bocchi) Strum(chord string) {}

var _ Guitarist = (*Bocchi)(nil)
