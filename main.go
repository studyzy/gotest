package main

import (
	"fmt"
	// "github.com/facebookgo/inject"
	// "github.com/studyzy/gotest/consensus"
	// _ "github.com/studyzy/gotest/core"
	// "github.com/studyzy/gotest/dag"
	"github.com/studyzy/gotest/db"
	"github.com/studyzy/gotest/statedb"
)

type Address [32]byte

func (a Address) Hex() string { return fmt.Sprintf("0x%x", a) }
func main() {
	fmt.Println("Hello World!")

	// var dag = new(dag.Dag)
	// var g inject.Graph
	// var m consensus.Mediator
	// g.Provide(
	// 	&inject.Object{Value: &m},
	// 	&inject.Object{Value: dag})
	// g.Populate()

	// fmt.Println(m.ReadSystemConfig("Count"))
	// var a = new(Address)
	// a[0] = byte('1')
	// fmt.Println(a.Hex())
	db, _ := db.NewMemDatabase()
	helper := statedb.NewStateHelper(db)
	helper.InitAcount("A", 100)
	helper.InitAcount("B", 100)
	helper.Transfer("A", "B", 10)
	valueA := helper.GetAcount("A")
	fmt.Printf("A amount:%d", valueA)
}
