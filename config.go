package gotest

// go get github.com/fjl/gencodec
//go:generate gencodec -type Config -field-override configMarshaling -formats toml -out gen_config.go

type Config struct {

	// Protocol options
	NetworkId uint64 // Network ID to use for selecting peers to connect to

	NoPruning bool

	// Light client options
	LightServ  int `toml:",omitempty"` // Maximum percentage of time allowed for serving LES requests
	LightPeers int `toml:",omitempty"` // Maximum number of LES client peers

	// Database options
	SkipBcVersionCheck bool `toml:"-"`
	DatabaseHandles    int  `toml:"-"`
	DatabaseCache      int
	TrieCache          int
}
