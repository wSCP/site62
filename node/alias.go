package node

import (
	"bazil.org/fuse"
	"github.com/thrisp/wSCP/xandle"
	"github.com/thrisp/wSCP/xandle/window"
)

type FindAliasFn func(string, Node, xandle.Xandle) (bool, Node)

type ListAliasFn func(Node, xandle.Xandle) []fuse.Dirent

type Aliaser interface {
	Aliased() bool
	Find(string, Node, xandle.Xandle) (bool, Node)
	List(Node, xandle.Xandle) []fuse.Dirent
}

type aliaser struct {
	aliased bool
	find    FindAliasFn
	list    ListAliasFn
}

func NewAliaser(a bool, f FindAliasFn, l ListAliasFn) Aliaser {
	return &aliaser{a, f, l}
}

func (a *aliaser) Aliased() bool {
	return a.aliased
}

func (a *aliaser) Find(requested string, n Node, x xandle.Xandle) (bool, Node) {
	return a.find(requested, n, x)
}

func (a *aliaser) List(n Node, x xandle.Xandle) []fuse.Dirent {
	return a.list(n, x)
}

var (
	NotAliased    Aliaser = NewAliaser(false, nil, nil)
	WindowAliaser Aliaser = NewAliaser(true, findWindowDir, listWindowDir)
)

func findWindowDir(requested string, n Node, x xandle.Xandle) (bool, Node) {
	if wid, exists := x.WindowExists(requested); exists {
		rd := n.Copy()
		rd.SetWindow(window.New(x.Conn(), wid, x.Root()))
		return true, rd
	}
	return false, nil
}

func listWindowDir(n Node, x xandle.Xandle) []fuse.Dirent {
	var ret []fuse.Dirent
	wl, err := x.ListWindows()
	if err == nil {
		for _, w := range wl {
			ret = append(ret, n.DirEntry(windowString(w)))
		}
	}
	return ret
}
