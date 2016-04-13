package node

import (
	"os"
	"strings"

	"golang.org/x/net/context"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/wSCP/site62/filo"
	"github.com/wSCP/site62/state"
)

// The Node interface is the primary set of interfaces for representing files and
// directories within a site62 filesystem.
type Node interface {
	Kind
	fs.Node
	NodeState
	NodePath
	NodeName
	Aliaser
	NodeMode
	NodeCopier
	NodeDirectory
	filo.Filo
	Tree
}

//
type NodeKind int

const (
	Unknown NodeKind = iota
	Directory
	File
	Fileio
	//Socket
)

// String returns a string for a NodeKind
func (n NodeKind) String() string {
	switch n {
	case Directory:
		return "directory"
	case File:
		return "file"
	case Fileio:
		return "file-io"
		//case Socket:
		//	return "socket"
	}
	return "unknown"
}

// StringNodeKind takes a string and returns a NodeKind
func StringNodeKind(s string) NodeKind {
	switch strings.ToLower(s) {
	case "directory":
		return Directory
	case "file":
		return File
	case "file-io":
		return Fileio
		//case "socket":
		//	return Socket
	}
	return Unknown
}

// The Kind interface orients a node as directory or file.
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

// Entry takes a string and return a fuse.Dirent.
func (k kind) Entry(name string) fuse.Dirent {
	switch k.k {
	case Directory:
		return entry(name, fuse.DT_Dir)
	case File:
		return entry(name, fuse.DT_File)
	case Fileio:
		return entry(name, fuse.DT_File)
		//case Socket:
		//	return entry(name, fuse.DT_Socket)
	}
	return entry(name, fuse.DT_Unknown)
}

type node struct {
	state state.State
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
		state:   nil,
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

// Attr satisfies the fuse/fs.Node interface for Node.
func (n *node) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Mode = n.Mode()
	return nil
}

type NodeState interface {
	Current() state.State
	SetState(state.State)
}

func (n *node) Current() state.State {
	return n.state
}

func (n *node) SetState(s state.State) {
	n.state = s
}

// NodePath is an interface for getting and setting a string path for a node.
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

// NodeName is an interface managing a node string name.
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

// NodeMode is an interface for setting & getting os.FileMode for a node.
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
	}
	n.mode = m
}

// NodeCopier is an interface providing methods to copy a node.
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
