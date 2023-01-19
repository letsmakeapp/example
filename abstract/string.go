package abstract

//go:generate mockgen -destination ./mock/string.go -package abstract_mock example/abstract StringLengthCounter

type StringLengthCounter interface {
	Count(s string) uint
}
