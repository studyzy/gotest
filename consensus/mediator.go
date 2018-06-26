package consensus

import (
	"github.com/studyzy/gotest/core"
)

type Mediator struct {
	dataHelper core.DataHelper `inject:""`
}

func (mediator Mediator) ConfirmUnit(unit core.Unit) {

}
func (mediator Mediator) ReadSystemConfig(key string) string {
	var helper = mediator.dataHelper
	return helper.GetSystemConfig(key)
}
