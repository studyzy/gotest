package main

import (
	"fmt"

	"github.com/facebookgo/inject"
	"github.com/studyzy/gotest/consensus"
	"github.com/studyzy/gotest/dag"
)

func main() {
	fmt.Println("Hello World!")

	var dag = new(dag.Dag)
	var g inject.Graph
	var m consensus.Mediator
	g.Provide(
		&inject.Object{Value: &m},
		&inject.Object{Value: dag})
	g.Populate()

	fmt.Println(m.ReadSystemConfig("Count"))
}
