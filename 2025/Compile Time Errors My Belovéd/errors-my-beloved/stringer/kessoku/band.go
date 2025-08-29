package kessoku

type BandMember int

const (
	Bocchi BandMember = iota
	Ryo
	Nijika
	Kita

	maxBand
)

var _ [0]struct{} = [maxBand - 4]struct{}{}
