package filesystem

import (
	"log"
	"os"
	"syscall"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/wSCP/site62/node"
	"github.com/wSCP/site62/state"
)

// Holds the primary file system data.
type FS struct {
	*log.Logger
	*settings
	Configuration
	*loop
	state state.State
	root  node.Node
}

// New returns an instance of FS with the provided Config.
func New(c ...Config) *FS {
	f := &FS{
		settings: newSettings(),
		loop:     newLoop(),
	}
	f.Configuration = newConfiguration(f, c...)
	return f
}

// Root satisfies the the fuse/fs FS interface.
func (f *FS) Root() (fs.Node, error) {
	if f.root == nil {
		f.root = f.RootFn(f.state, f.RootPath)
		for _, m := range f.Mountable {
			err := Attach(m, f.root)
			if err != nil {
				f.Printf("unable to attach mounts %s to any %s in filesystem", m.Kind(), m.At())
			}
		}
	}
	return f.root, nil
}

//
func (f *FS) Mount() error {
	if !f.Configured() {
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

func (f *FS) signalHandler(sig os.Signal) {
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
