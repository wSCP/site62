package filesystem

import (
	"os"
	"syscall"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/thrisp/wSCP/site62/node"
	"github.com/thrisp/wSCP/xandle"
)

// Holds the primary file system data.
type FS struct {
	*Configuration
	*loop
	xandle.Xandle
	root node.Node
}

//
func New(c ...ConfigureFn) *FS {
	return &FS{
		Configuration: NewConfiguration(c...),
		loop:          newLoop(),
	}
}

// Root satisfies the the fuse/fs FS interface.
func (f *FS) Root() (fs.Node, error) {
	if f.root == nil {
		f.root = f.RootFn(f.Xandle, f.RootPath)
		for _, n := range f.mounts {
			f.root.SetTail(n.Block())
		}
	}
	return f.root, nil
}

func (f *FS) Mount() error {
	if !f.configured {
		if confErr := f.Configure(); confErr != nil {
			return confErr
		}
	}

	c, mountErr := fuse.Mount(f.RootPath)
	if mountErr != nil {
		return mountErr
	}
	defer c.Close()

	looping(f)

	if serveErr := fs.Serve(c, f); serveErr != nil {
		return serveErr
	}

	<-c.Ready
	if errd := c.MountError; errd != nil {
		return errd
	}

	return nil
}

func (f *FS) SignalHandler(sig os.Signal) {
	switch sig {
	case syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM:
		fuse.Unmount(f.RootPath)
		os.Exit(0)
	}
}

// Root satisfies the the fuse/fs FS interface.
func (f *FS) Destroy() {
	if f.root != nil {
		f.root = nil
	}
	f = nil
}
