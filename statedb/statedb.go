package statedb

import (
	"github.com/studyzy/gotest/db"
	"strconv"
)

type Statehelper struct {
	Db db.Database
}

func NewStateHelper(db db.Database) *Statehelper {
	return &Statehelper{
		Db: db,
	}
}
func (s *Statehelper) Transfer(aname, bname string, amount int) int {
	valueA, _ := strconv.Atoi(string(s.Db.Get([]byte(aname))))
	valueB, _ := strconv.Atoi(string(s.Db.Get([]byte(bname))))
	valueA += amount
	valueB -= amount
	s.Db.Put([]byte(aname), []byte(strconv.Itoa(valueA)))
	s.Db.Put([]byte(bname), []byte(strconv.Itoa(valueB)))
	return valueA
}
func (s *Statehelper) InitAcount(name string, amount int) {
	s.Db.Put([]byte(name), []byte(strconv.Itoa(amount)))
}
func (s *Statehelper) GetAcount(name string) int {
	value, _ := strconv.Atoi(string(s.Db.Get([]byte(name))))
	return value
}
