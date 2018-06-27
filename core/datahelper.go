package core

import (
	"fmt"
)

//Data interface
type DataHelper interface {
	GetSystemConfig(key string) string
	WriteUnit(unit Unit)
}

func init() {
	fmt.Println("core-init() come here.")
}
