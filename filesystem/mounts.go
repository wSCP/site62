package filesystem

import (
	"errors"

	"github.com/wSCP/site62/node"
)

type Mounts interface {
	Kind() string
	At() string
	Block() node.Node
}

type mounts struct {
	kind  string
	at    string
	block node.Node
}

func (f *FS) NewMounts(k, at string, b node.Node) {
	m := &mounts{k, at, b}
	f.Mountable = append(f.Mountable, m)
}

func (m *mounts) Kind() string {
	return m.kind
}

func (m *mounts) At() string {
	return m.at
}

func (m *mounts) Block() node.Node {
	return m.block
}

var Unattachable = errors.New("unattachable")

func attach(m Mounts, to node.Node) error {
	if m.At() == to.Name() {
		to.SetTail(m.Block())
		return nil
	}
	for _, t := range to.Tail() {
		attach(m, t)
	}
	return Unattachable
}

func Attach(m Mounts, to node.Node) error {
	return attach(m, to)
}
