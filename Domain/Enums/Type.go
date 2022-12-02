package Enums

//go:generate stringer -type=Type

type Type byte

const (
	TypeTest Type = iota + 1
)
