package dag

import (
	"github.com/studyzy/gotest/core"
)

type Dag struct{}

func (dag *Dag) GetSystemConfig(key string) string {
	return "Dag retrieve value for key: " + key
}
func (dag *Dag) WriteUnit(unit core.Unit) {

}
