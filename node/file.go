package node

import (
	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/thrisp/wSCP/site62/filo"
	"golang.org/x/net/context"
)

//
type NodeFile interface {
	InitializeFile(path string, head Node) NodeFile
	Node
	filo.Filo
}

type nodeFile struct {
	Node
	filo.Filo
}

// NewNodeFile
func NewNodeFile(n Node) NodeFile {
	return &nodeFile{n, filo.Get(n.Nio())}
}

// Initialize prepares the NodeFile from the given string path and head Node.
func (n *nodeFile) InitializeFile(path string, head Node) NodeFile {
	rf := n
	rf.SetPath(path)
	rf.SetHead(head)
	rf.SetXandle(head.Xandle())
	rf.SetWindow(head.Window())
	rf.Filo.Init(rf.Xandle(), rf.Window())
	return rf
}

// Attr satisfies the fuse/fs Node interface for a nodeFile instance.
func (n *nodeFile) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Mode = n.Mode()
	if n.Is() != Fileio {
		a.Size = n.Size()
	}
	return nil
}

// Open satisfies the fuse/fs NodeOpener interface for a nodeFile instance.
func (n *nodeFile) Open(ctx context.Context, req *fuse.OpenRequest, resp *fuse.OpenResponse) (fs.Handle, error) {
	if n.Is() == Fileio {
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
