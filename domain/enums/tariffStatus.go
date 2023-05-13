package enums

//go:generate stringer -type=TariffStatus -trimprefix=TariffStatus -output=tariffStatus_string.go

type TariffStatus byte

const (
	TariffStatusEnabled TariffStatus = iota + 1
	TariffStatusDisabled
)
