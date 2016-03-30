package node

import (
	"os"

	"github.com/wSCP/site62/filo"
	"github.com/wSCP/xandle"
)

func NewRoot(x xandle.Xandle, path string) Node {
	r := &Root
	r.SetXandle(x)
	r.SetPath(path)
	return r
}

var Root node = node{
	Header:  &EmptyHeader,
	Kind:    kind{Directory},
	name:    "/",
	mode:    os.ModeDir | 0700,
	Aliaser: NotAliased,
	Tree: NewTree(nil, BaseDir("root", 0700,
		BaseDir("geometry", 0700,
			BaseFile("width", 0400, "root_width"),
			BaseFile("height", 0400, "root_height"),
		)),
	),
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
