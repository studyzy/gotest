package dag

import (
	"fmt"
	"github.com/studyzy/gotest/core"
)

type Dag struct{}

func (dag *Dag) GetSystemConfig(key string) string {
	return "Dag retrieve value for key: " + key
}
func (dag *Dag) WriteUnit(unit core.Unit) {

}
func init() {
	fmt.Println("dag-init() come here.")
}
