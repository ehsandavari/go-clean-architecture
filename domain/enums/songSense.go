package enums

//go:generate stringer -type=SongSense -trimprefix=SongSense -output=songSense_string.go

type SongSense byte

const (
	SongSenseHappyPositive SongSense = iota + 1
	SongSenseSadNegative
)
