package enums

//go:generate stringer -type=SongGenre -trimprefix=SongGenre -output=songGenre_string.go

type SongGenre byte

const (
	SongGenrePop SongGenre = iota + 1
	SongGenreRap
	SongGenreRock
	SongGenreHipHop
	SongGenreTraditional
)
