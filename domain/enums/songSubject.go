package enums

//go:generate stringer -type=SongSubject -trimprefix=Subject -output=songSubject_string.go

type SongSubject byte

const (
	SongSubjectRomantic SongSubject = iota + 1
	SongSubjectSocial
	SongSubjectEpic
	SongSubjectReligious
)
