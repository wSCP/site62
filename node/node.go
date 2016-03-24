package node

import (
	"os"
	"strings"

	"golang.org/x/net/context"

	"bazil.org/fuse"
	"github.com/thrisp/wSCP/site62/filo"
)

//
type Node interface {
	Header
	Kind
	Attr(ctx context.Context, a *fuse.Attr) error
	NodePath
	NodeName
	NodeMode
	NodeCopier
	NodeDirectory
	filo.Filo
	Aliaser
	Tree
}

//
type NodeKind int

//
func (n NodeKind) String() string {
	switch n {
	case Directory:
		return "directory"
	case File:
		return "file"
	case Fileio:
		return "file-io"
	case Socket:
		return "socket"
	}
	return "unknown"
}

//
func stringNodeKind(s string) NodeKind {
	switch strings.ToLower(s) {
	case "directory":
		return Directory
	case "file":
		return File
	case "file-io":
		return Fileio
	case "socket":
		return Socket
	}
	return Unknown
}

const (
	Unknown NodeKind = iota
	Directory
	File
	Fileio
	Socket
)

//
type Kind interface {
	Is() NodeKind
	Entry(name string) fuse.Dirent
}

type kind struct {
	k NodeKind
}

//
func (k kind) Is() NodeKind {
	return k.k
}

//
func (k kind) Entry(name string) fuse.Dirent {
	switch k.k {
	case Directory:
		return entry(name, fuse.DT_Dir)
	case File:
		return entry(name, fuse.DT_File)
	case Fileio:
		return entry(name, fuse.DT_File)
	case Socket:
		return entry(name, fuse.DT_Socket)
	}
	return entry(name, fuse.DT_Unknown)
}

type node struct {
	Header
	Kind
	path, name string
	mode       os.FileMode
	filo.Filo
	Aliaser
	Tree
}

// New returns a Node instance fromt the provided parameters..
func New(k NodeKind, path, name string, mode os.FileMode, f filo.Filo, a Aliaser, head Node, tail ...Node) Node {
	n := &node{
		Header:  &EmptyHeader,
		Kind:    kind{k},
		path:    path,
		name:    name,
		Filo:    f,
		Aliaser: a,
		Tree:    NewTree(head, tail...),
	}
	n.SetMode(mode)
	return n
}

// Attr satisfies the fuse/fs Node interface for Node.
func (n *node) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Mode = n.Mode()
	return nil
}

// An interface for getting and setting a string path for a node.
type NodePath interface {
	Path() string
	SetPath(string)
}

// Path returns the path containing the node.
func (n *node) Path() string {
	return n.path
}

// SetPath sets the path containing the node.
func (n *node) SetPath(p string) {
	n.path = p
}

// An interface managing a node name.
type NodeName interface {
	Name() string
	SetName(string)
}

// Name returns the name of the node.
func (n *node) Name() string {
	return n.name
}

// SetName sets the node name by the provided string.
func (n *node) SetName(nm string) {
	n.name = nm
}

// An interface for setting & getting a node mode.
type NodeMode interface {
	Mode() os.FileMode
	SetMode(os.FileMode)
}

// Mode returns a node's mode as os.Filemode
func (n *node) Mode() os.FileMode {
	return n.mode
}

// SetMode takes os.FileMode input to set the node mode. os.ModeDir is set for
// directories here.
func (n *node) SetMode(m os.FileMode) {
	switch n.Is() {
	case Directory:
		m = (os.ModeDir | m)
	case Socket:
		m = (os.ModeSocket | m)
	}
	n.mode = m
}

// An interface providing methods to copy a node.
type NodeCopier interface {
	Copy() Node
}

// Copy returns a copy of the node. Notable default is the replacement of the
// Aliaser with one that provides no alias.
func (n *node) Copy() Node {
	var nn node = *n
	nn.Aliaser = NotAliased
	return &nn
}
