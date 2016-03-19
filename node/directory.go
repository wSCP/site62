package node

import (
	"path/filepath"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"golang.org/x/net/context"
)

type NodeDirectory interface {
	InitializeDir(path string, head Node) Node
	Lookup(ctx context.Context, req *fuse.LookupRequest, resp *fuse.LookupResponse) (fs.Node, error)
	ReadDirAll(ctx context.Context) ([]fuse.Dirent, error)
}

func (n *node) InitializeDir(path string, head Node) Node {
	n.SetPath(path)
	n.SetHead(head)
	n.SetXandle(head.Xandle())
	n.SetWindow(head.Window())
	return n
}

// Lookup satisfies the fuse/fs NodeLookuper interface for a node.
func (n *node) Lookup(ctx context.Context, req *fuse.LookupRequest, resp *fuse.LookupResponse) (fs.Node, error) {
	var requested string = req.Name
	p, nm := n.Path(), n.Name()
	fp, _ := filepath.Abs(filepath.Join(p, "/", nm))
	for _, nn := range n.Tail() {
		if nn.Name() == requested {
			switch nn.Is() {
			case Directory:
				d := nn.InitializeDir(fp, n)
				return d, nil
			case File, Fileio:
				f := NewNodeFile(nn)
				f.InitializeFile(fp, n)
				return f, nil
			}
		}
		if nn.Aliased() {
			if exists, ad := nn.Find(requested, nn, n.Xandle()); exists {
				switch ad.Is() {
				case Directory:
					ad.InitializeDir(fp, n)
					return ad, nil
				case File, Fileio:
					af := NewNodeFile(ad)
					af.InitializeFile(fp, n)
					return af, nil
				}
			}
		}

	}
	return nil, fuse.ENOENT
}

func dots() []fuse.Dirent {
	var ret []fuse.Dirent
	ret = append(ret, entry(".", fuse.DT_Dir))
	ret = append(ret, entry("..", fuse.DT_Dir))
	return ret
}

// ReadDirAll satisfies the fuse/fs HandleReadDirAller interface for a node.
func (n *node) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	var ret []fuse.Dirent
	ret = append(ret, dots()...)
	for _, nn := range n.Tail() {
		if nn.Aliased() {
			ret = append(ret, nn.List(nn, n.Xandle())...)
		} else {
			ret = append(ret, nn.DirEntry(nn.Name()))
		}
	}
	return ret, nil
}
