package checks

import (
	"validity/pkg/wdb"
)

type Client struct {
	WDBClient wdb.Client
	ControlList string
}

var (
	Valid   = true
	Invalid = false
)