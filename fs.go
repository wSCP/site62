package main

import (
	"bazil.org/fuse/fs"
	"github.com/thrisp/wSCP/site62/node"
	"github.com/thrisp/wSCP/xandle"
)

// Holds the primary data for operating site62.
type FS struct {
	xandle.Xandle
	mount string
	root  node.Node
}

//
func New(x xandle.Xandle, mount string) *FS {
	return &FS{
		Xandle: x,
		mount:  mount,
	}
}

// Root satisfies the the fuse/fs FS interface.
func (f *FS) Root() (fs.Node, error) {
	if f.root == nil {
		f.root = node.NewRoot(f.Xandle, f.mount)
	}
	return f.root, nil
}

// Root satisfies the the fuse/fs FS interface.
func (f *FS) Destroy() {
	if f.root != nil {
		f.root = nil
	}
	f = nil
}
