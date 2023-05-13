package enums

//go:generate stringer -type=SongStatus -trimprefix=SongStatus -output=songStatus_string.go

type SongStatus byte

const (
	SongStatusPending SongStatus = iota + 1
	SongStatusVerified
	SongStatusRejected
	SongStatusSold
)
