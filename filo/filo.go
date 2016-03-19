package filo

import (
	"sync"
	"syscall"

	"bazil.org/fuse"
	"golang.org/x/net/context"

	"github.com/thrisp/wSCP/xandle"
	"github.com/thrisp/wSCP/xandle/window"
)

type Filo interface {
	Key() string
	Init(xandle.Xandle, window.Window)
	Read(ctx context.Context, req *fuse.ReadRequest, resp *fuse.ReadResponse) error
	Write(ctx context.Context, req *fuse.WriteRequest, resp *fuse.WriteResponse) error
	Size() uint64
}

type filo struct {
	xandle.Xandle
	Window window.Window
	key    string
	read   FileRead
	write  FileWrite
	sync.RWMutex
}

// New creates a new Filo with key, read, and write functions.
func New(key string, read FileRead, write FileWrite) Filo {
	return &filo{
		key:   key,
		read:  read,
		write: write,
	}
}

// Key returns the filo key as a string.
func (f *filo) Key() string {
	return f.key
}

// Init providse A Xandle and a Window to the filo.
func (f *filo) Init(x xandle.Xandle, w window.Window) {
	f.Xandle = x
	f.Window = w
}

// Read allows filo to satisfy the fuse/fs HandleReader interface.
func (f *filo) Read(ctx context.Context, req *fuse.ReadRequest, resp *fuse.ReadResponse) error {
	if f.read == nil {
		return syscall.EACCES
	}
	f.RLock()
	defer f.RUnlock()
	res := f.read(f.Window, f.Xandle)
	resp.Data = res.Bytes()
	return nil
}

// Read allows filo to satisfy the fuse/fs HandleWriter interface for Filo
func (f *filo) Write(ctx context.Context, req *fuse.WriteRequest, resp *fuse.WriteResponse) error {
	if f.write == nil {
		return syscall.EACCES
	}
	f.Lock()
	defer f.Unlock()
	wrote, err := f.write(f.Window, f.Xandle, req.Data)
	resp.Size = wrote
	return err
}

// Size return a uint64 measure of the filo size in bytes.
func (f *filo) Size() uint64 {
	if f.read != nil {
		res := f.read(f.Window, f.Xandle)
		return uint64(res.Len())
	}
	return 0
}
