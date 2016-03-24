package node

import (
	"path/filepath"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"golang.org/x/net/context"
)

//
type NodeDirectory interface {
	InitializeDir(path string, head Node)
	Lookup(ctx context.Context, req *fuse.LookupRequest, resp *fuse.LookupResponse) (fs.Node, error)
	ReadDirAll(ctx context.Context) ([]fuse.Dirent, error)
}

//
func (n *node) InitializeDir(path string, head Node) {
	n.SetPath(path)
	n.SetHead(head)
	n.SetXandle(head.Xandle())
	n.SetWindow(head.Window())
}

// Lookup satisfies the fuse/fs NodeLookuper interface for a node.
func (n *node) Lookup(ctx context.Context, req *fuse.LookupRequest, resp *fuse.LookupResponse) (fs.Node, error) {
	var requested string = req.Name
	p, nm := n.Path(), n.Name()
	fp, _ := filepath.Abs(filepath.Join(p, "/", nm))
	tail := n.Tail()
	if len(tail) > 0 {
		for _, nn := range tail {
			if nn.Aliased() {
				if exists, ad := nn.Find(requested, n, nn); exists {
					switch ad.Is() {
					case Directory:
						ad.SetPath(fp)
						return ad, nil
					case File, Fileio:
						af := NewNodeFile(ad)
						af.InitializeFile(fp, n)
						return af, nil
					}
				}
			}
			if nn.Name() == requested {
				switch nn.Is() {
				case Directory:
					nn.InitializeDir(fp, n)
					return nn, nil
				case File, Fileio, Socket:
					f := NewNodeFile(nn)
					f.InitializeFile(fp, n)
					return f, nil
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
	tail := n.Tail()
	if len(tail) > 0 {
		for _, nn := range n.Tail() {
			if nn.Aliased() {
				ret = append(ret, nn.List(nn, n.Xandle())...)
			} else {
				ret = append(ret, nn.Entry(nn.Name()))
			}
		}
	}
	return ret, nil
}

//func (n *node) Create(ctx context.Context, req *fuse.CreateRequest, resp *fuse.CreateResponse) (fs.Node, fs.Handle, error) {
//	spew.Dump(req)
//	return nil, nil, errors.New("NOPE")
//}
