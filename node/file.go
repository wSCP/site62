package node

import (
	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"golang.org/x/net/context"
)

//
type NodeFile interface {
	InitializeFile(path string, head Node)
	Node
}

type nodeFile struct {
	Node
}

// NewNodeFile returns a NodeFile wrapping the provided Node.
func NewNodeFile(n Node) NodeFile {
	return &nodeFile{n}
}

// Initialize prepares the NodeFile from the given string path and head Node.
func (n *nodeFile) InitializeFile(path string, head Node) {
	n.SetPath(path)
	n.SetHead(head)
	n.SetXandle(head.Xandle())
	n.SetWindow(head.Window())
	n.Init(n.CurrentHeader())
}

// Attr satisfies the fuse/fs Node interface for NodeFile.
func (n *nodeFile) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Mode = n.Mode()
	if n.Is() != Fileio {
		a.Size = n.Size()
	}
	return nil
}

// Open satisfies the fuse/fs NodeOpener interface for a nodeFile instance.
func (n *nodeFile) Open(ctx context.Context, req *fuse.OpenRequest, resp *fuse.OpenResponse) (fs.Handle, error) {
	switch n.Is() {
	case Fileio:
		resp.Flags = resp.Flags << fuse.OpenDirectIO
	}
	resp.Flags = resp.Flags << fuse.OpenNonSeekable
	return n, nil
}

// Remove satisfies the fuse/fs NodeRemover interface for File.
//func (f *File) Remove(ctx context.Context, req *fuse.RemoveRequest) error {
//	f = nil
//	return nil
//}

// Read satisfies the fuse/fs HandleReleaser interface for File.
//func (f *File) Release(ctx context.Context, req *fuse.ReleaseRequest) error {
//	return nil
//}
