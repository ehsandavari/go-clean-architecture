package enums

//go:generate stringer -type=SongLanguage -trimprefix=SongLanguage -output=songLanguage_string.go

type SongLanguage byte

const (
	SongLanguagePersian SongLanguage = iota + 1
	SongLanguageEnglish
	SongLanguageArabic
	SongLanguageTurkish
	SongLanguageKurdish
	SongLanguageOther
)
