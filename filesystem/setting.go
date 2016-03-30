package filesystem

import (
	"github.com/wSCP/site62/node"
	"github.com/wSCP/xandle"
)

type settings struct {
	verbose   bool
	RootFn    func(xandle.Xandle, string) node.Node
	RootPath  string
	Mountable []Mounts
}

func newSettings() *settings {
	return &settings{
		Mountable: make([]Mounts, 0),
	}
}
