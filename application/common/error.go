package common

//go:generate stringer -type=tError -trimprefix=Error -output=error_string.go

type tError uint

func (r tError) Error() string {
	return r.String()
}

func (r tError) Code() uint {
	return uint(r)
}

const (
	ErrorInternal tError = iota + 1000
	ErrorOrderNotFound
)
