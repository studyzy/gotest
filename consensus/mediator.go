package consensus

import (
	"fmt"
	"github.com/studyzy/gotest/core"
)

func init() {
	fmt.Println("consensus-init() come here.")
}

type Mediator struct {
	DataHelper core.DataHelper `inject:""`
}

func (mediator Mediator) ConfirmUnit(unit core.Unit) {

}
func (mediator Mediator) ReadSystemConfig(key string) string {
	var helper = mediator.DataHelper
	return helper.GetSystemConfig(key)
}
