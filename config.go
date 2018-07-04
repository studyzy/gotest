package main

type DatabaseConfig struct {
	Host string `mapstructure:"hostname"`
	Port string
	User string `mapstructure:"username"`
	Pass string `mapstructure:"password"`
}
type FileStruct struct {
	Path PathStruct
	Ext string
}
type PathStruct struct {
	FullPath string
	Folder string
}
type OutputConfig struct {
	File FileStruct
}

type Config struct {
	Db  DatabaseConfig `mapstructure:"database"`
	Out OutputConfig   `mapstructure:"output"`
}
