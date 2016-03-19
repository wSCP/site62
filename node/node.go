package node

import (
	"os"
	"strings"

	"golang.org/x/net/context"

	"bazil.org/fuse"
)

type Node interface {
	Header
	Kind
	Attr(ctx context.Context, a *fuse.Attr) error
	Path() string
	SetPath(string)
	Name() string
	SetName(string)
	Mode() os.FileMode
	SetMode(os.FileMode)
	Nio() string
	NodeCopier
	NodeDirectory
	Aliaser
	Tree
}

type NodeKind int

func (n NodeKind) String() string {
	switch n {
	case Directory:
		return "directory"
	case File:
		return "file"
	case Fileio:
		return "file-io"
	}
	return "unknown"
}

func stringNodeKind(s string) NodeKind {
	switch strings.ToLower(s) {
	case "directory":
		return Directory
	case "file":
		return File
	case "file-io":
		return Fileio
	}
	return Unknown
}

const (
	Unknown NodeKind = iota
	Directory
	File
	Fileio
)

type Kind interface {
	Is() NodeKind
	DirEntry(name string) fuse.Dirent
}

type kind struct {
	k NodeKind
}

func (k kind) Is() NodeKind {
	return k.k
}

func (k kind) DirEntry(name string) fuse.Dirent {
	switch k.k {
	case Directory:
		return entry(name, fuse.DT_Dir)
	case File:
		return entry(name, fuse.DT_File)
	case Fileio:
		return entry(name, fuse.DT_File)
	}
	return entry(name, fuse.DT_Unknown)
}

type node struct {
	Header
	Kind
	path, name, nio string
	mode            os.FileMode
	Aliaser
	Tree
}

func newNode(nk NodeKind, path, name, nio string, mode os.FileMode, a Aliaser, h Node, t ...Node) Node {
	n := &node{
		Header:  NewHeader(nil, nil),
		Kind:    kind{nk},
		path:    path,
		name:    name,
		nio:     nio,
		Aliaser: a,
		Tree:    NewTree(h, t...),
	}
	n.SetMode(mode)
	return n
}

// Attr satisfies the fuse/fs Node interface for Node.
func (n *node) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Mode = n.Mode()
	return nil
}

//
func (n *node) Path() string {
	return n.path
}

//
func (n *node) SetPath(p string) {
	n.path = p
}

//
func (n *node) Name() string {
	return n.name
}

//
func (n *node) SetName(nm string) {
	n.name = nm
}

//
func (n *node) Mode() os.FileMode {
	return n.mode
}

//
func (n *node) SetMode(m os.FileMode) {
	if n.Is() == Directory {
		m = (os.ModeDir | m)
	}
	n.mode = m
}

//
func (n *node) Nio() string {
	return n.nio
}

type NodeCopier interface {
	Copy() Node
}

func (n *node) Copy() Node {
	var nn node = *n
	nn.Aliaser = NotAliased
	return &nn
}
