package db

type Database interface {
	Get(key []byte) []byte
	Put(key []byte, value []byte) error
}
