package enums

//go:generate stringer -type=Type -trimprefix=Type

type Type byte

const (
	TypeTest Type = iota + 1
)
