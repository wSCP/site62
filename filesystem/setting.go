package filesystem

import (
	"github.com/wSCP/site62/node"
	"github.com/wSCP/site62/state"
)

type settings struct {
	verbose   bool
	RootFn    func(state.State, string) node.Node
	RootPath  string
	Mountable []Mounts
}

func newSettings() *settings {
	return &settings{
		Mountable: make([]Mounts, 0),
	}
}
