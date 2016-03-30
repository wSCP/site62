package node

import (
	"bazil.org/fuse"
	"github.com/wSCP/xandle"
)

type AliasFindFn func(string, Node, Node) (bool, Node)

type AliasListFn func(Node, xandle.Xandle) []fuse.Dirent

type Aliaser interface {
	AliasKey() string
	Aliased() bool
	Find(string, Node, Node) (bool, Node)
	List(Node, xandle.Xandle) []fuse.Dirent
}

type aliaser struct {
	key     string
	aliased bool
	find    AliasFindFn
	list    AliasListFn
}

func NewAliaser(k string, a bool, f AliasFindFn, l AliasListFn) Aliaser {
	return &aliaser{k, a, f, l}
}

func (a *aliaser) AliasKey() string {
	return a.key
}

func (a *aliaser) Aliased() bool {
	return a.aliased
}

func (a *aliaser) Find(requested string, h Node, n Node) (bool, Node) {
	return a.find(requested, h, n)
}

func (a *aliaser) List(n Node, x xandle.Xandle) []fuse.Dirent {
	return a.list(n, x)
}

type aliasers map[string]Aliaser

func (as aliasers) SetAliasers(a ...Aliaser) {
	for _, al := range a {
		as[al.AliasKey()] = al
	}
}

func (as aliasers) GetAliaser(key string) Aliaser {
	if ret, exists := as[key]; exists {
		return ret
	}
	return NotAliased
}

var Aliasers aliasers

func init() {
	Aliasers = make(aliasers)
	Aliasers.SetAliasers(NotAliased)
}

var NotAliased Aliaser = NewAliaser("not_aliased", false, nil, nil)
