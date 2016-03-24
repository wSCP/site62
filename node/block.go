package node

import (
	"os"

	"github.com/thrisp/wSCP/site62/filo"
	"github.com/thrisp/wSCP/xandle"
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
	Tree: NewTree(nil, baseDir("root", 0700,
		baseDir("geometry", 0700,
			baseFile("width", 0400, "root_width"),
			baseFile("height", 0400, "root_height"),
		)),
		baseSocket("event", 0777, ""),
	),
}

var ClientsBlock node = node{
	Header:  &EmptyHeader,
	Kind:    kind{Directory},
	name:    "clients",
	mode:    os.ModeDir | 0700,
	Aliaser: NotAliased,
	Tree: NewTree(nil, baseAliasedDir("0x*", 0700, WindowAliaser,
		baseDir("border", 0700,
			baseFile("color", 0200, "border_color"),
			baseFile("width", 0600, "border_width"),
		),
		baseDir("geometry", 0700,
			baseFile("area", 0600, "area"),
			baseFile("width", 0600, "width"),
			baseFile("height", 0600, "height"),
			baseFile("x", 0600, "X"),
			baseFile("y", 0600, "Y"),
		),
		baseDir("state", 0700,
			baseFile("map_state", 0600, "mapstate"),
			baseFile("mapped", 0600, "mapped"),
			//baseFile("focus", 0600, ""),
			//baseFile("ignored", 0600, ""),
			//baseFile("stack", 0200, ""),
		),
		baseFile("instance", 0400, "instance"),
		baseFile("class", 0400, "class"),
	),
		baseFile("focus", 0600, ""),
	),
}

var MonitorsBlock node = node{
	Header:  &EmptyHeader,
	Kind:    kind{Directory},
	name:    "monitors",
	mode:    os.ModeDir | 0700,
	Aliaser: NotAliased,
	Tree: NewTree(nil, baseAliasedDir("1x*", 0700, MonitorAliaser,
		baseDir("geometry", 0700,
			baseFile("area", 0600, ""),
			baseFile("width", 0600, ""),
			baseFile("height", 0600, ""),
			baseFile("x", 0600, ""),
			baseFile("y", 0600, ""),
		),
		baseFile("root", 0400, ""),
		baseFile("wired", 0400, ""),
		baseFile("focus", 0400, ""),
	),
		baseFile("focus", 0600, ""),
		baseFile("primary", 0600, ""),
	),
}

func BaseNode(k NodeKind, nm, flo string, a Aliaser, m os.FileMode, t ...Node) Node {
	return New(k, "", nm, m, filo.Get(flo), a, nil, t...)
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

func baseFileIO(name string, mode os.FileMode, flo string) Node {
	return BaseNode(Fileio, name, flo, NotAliased, mode)
}

func baseSocket(name string, mode os.FileMode, flo string) Node {
	return BaseNode(Socket, name, flo, NotAliased, mode)
}
