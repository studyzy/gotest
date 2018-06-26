package core

//Data interface
type DataHelper interface {
	GetSystemConfig(key string) string
	WriteUnit(unit Unit)
}
