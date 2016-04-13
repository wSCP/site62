package node

import (
	"os"

	"github.com/wSCP/site62/filo"
	"github.com/wSCP/site62/state"
)

func NewRoot(s state.State, path string) Node {
	r := &Root
	r.SetState(s)
	r.SetPath(path)
	return r
}

var Root node = node{
	Kind:    kind{Directory},
	name:    "/",
	mode:    os.ModeDir | 0700,
	Aliaser: NotAliased,
	Tree:    NewTree(nil),
}

func BaseNode(k NodeKind, nm, flo string, a Aliaser, m os.FileMode, t ...Node) Node {
	return New(k, "", nm, m, filo.Get(flo), a, nil, t...)
}

func BaseDir(name string, mode os.FileMode, t ...Node) Node {
	return BaseNode(Directory, name, "", NotAliased, mode, t...)
}

func BaseAliasedDir(name string, mode os.FileMode, a Aliaser, t ...Node) Node {
	return BaseNode(Directory, name, "", a, mode, t...)
}

func BaseFile(name string, mode os.FileMode, nio string) Node {
	return BaseNode(File, name, nio, NotAliased, mode)
}

//func BaseFileIO(name string, mode os.FileMode, flo string) Node {
//	return BaseNode(Fileio, name, flo, NotAliased, mode)
//}

//func baseSocket(name string, mode os.FileMode, flo string) Node {
//	return BaseNode(Socket, name, flo, NotAliased, mode)
//}
