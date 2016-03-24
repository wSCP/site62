package filesystem

import (
	"github.com/thrisp/wSCP/site62/node"
)

type MountsKind int

const (
	unknown MountsKind = iota
	monitors
	tags
	clients
)

func (mk MountsKind) String() string {
	switch mk {
	case monitors:
		return "MONITORS"
	case tags:
		return "TAGS"
	case clients:
		return "CLIENTS"
	}
	return "UNKNOWN"
}

type Mounts interface {
	Kind() string
	Block() node.Node
}

type mounts struct {
	kind  MountsKind
	block node.Node
}

func (f *FS) NewMounts(mk MountsKind, b node.Node) {
	m := &mounts{mk, b}
	f.mounts = append(f.mounts, m)
}

func (m *mounts) Kind() string {
	return m.kind.String()
}

func (m *mounts) Block() node.Node {
	return m.block
}
