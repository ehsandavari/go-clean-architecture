package common

//go:generate stringer -type=Error -trimprefix=Error

type Error byte

func (r Error) Error() string {
	return r.String()
}

const (
	ErrorOrderNotFound Error = iota + 1
)
