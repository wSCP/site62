package node

import (
	"os"

	"github.com/thrisp/wSCP/xandle"
)

func NewRoot(x xandle.Xandle, path string) Node {
	r := &Root
	r.SetXandle(x)
	r.SetPath(path)
	return r
}

var Root node = node{
	Header: NewHeader(nil, nil),
	Kind:   kind{Directory},
	name:   "/",
	mode:   os.ModeDir | 0700,
	Tree: NewTree(nil, baseDir("root", 0700,
		baseDir("geometry", 0700,
			baseFile("width", 0400, "root_width"),
			baseFile("height", 0400, "root_height"),
		)),
		baseDir("nodes", 0700,
			baseAliasedDir("0x*", 0700, WindowAliaser,
				baseDir("border", 0700,
					baseFile("color", 0200, "border_color"),
					baseFile("width", 0600, "border_width"),
				),
				baseDir("geometry", 0700,
					baseFile("width", 0600, "width"),
					baseFile("height", 0600, "height"),
					baseFile("x", 0600, "X"),
					baseFile("y", 0600, "Y"),
				),
				baseFile("mapped", 0600, ""),
				baseFile("ignored", 0600, ""),
				baseFile("stack", 0200, ""),
				baseFile("instance", 0400, "instance"),
				baseFile("class", 0400, "class"),
			),
			baseFile("focused", 0600, ""),
		),
		baseFileIO("event", 0400, ""),
	),
}

func BaseNode(k NodeKind, nm, i string, a Aliaser, m os.FileMode, t ...Node) Node {
	n := &node{
		Header:  NewHeader(nil, nil),
		Kind:    kind{k},
		name:    nm,
		nio:     i,
		Aliaser: a,
		Tree:    NewTree(nil, t...),
	}
	n.SetMode(m)
	return n
}

func baseDir(name string, mode os.FileMode, t ...Node) Node {
	return BaseNode(Directory, name, "", NotAliased, mode, t...)
}

func baseAliasedDir(name string, mode os.FileMode, a Aliaser, t ...Node) Node {
	return BaseNode(Directory, name, "", a, mode, t...)
}

func baseFile(name string, mode os.FileMode, nio string) Node {
	return BaseNode(File, name, nio, NotAliased, mode)
}

func baseFileIO(name string, mode os.FileMode, nio string) Node {
	return BaseNode(Fileio, name, nio, NotAliased, mode)
}
