package filo

import (
	"sync"
	"syscall"

	"bazil.org/fuse"
	"github.com/wSCP/xandle"
	"github.com/wSCP/xandle/monitor"
	"github.com/wSCP/xandle/window"
	"golang.org/x/net/context"
)

//
type Filo interface {
	Key() string
	Init(Header)
	Read(ctx context.Context, req *fuse.ReadRequest, resp *fuse.ReadResponse) error
	Write(ctx context.Context, req *fuse.WriteRequest, resp *fuse.WriteResponse) error
	Size() uint64
}

// A local interface mirroring node.Header.
type Header interface {
	Xandle() xandle.Xandle
	SetXandle(xandle.Xandle)
	Window() window.Window
	SetWindow(window.Window)
	Monitor() monitor.Monitor
	SetMonitor(monitor.Monitor)
}

type filo struct {
	h     Header
	key   string
	read  FileRead
	write FileWrite
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
func (f *filo) Init(h Header) {
	f.h = h
}

// Read allows filo to satisfy the fuse/fs HandleReader interface.
func (f *filo) Read(ctx context.Context, req *fuse.ReadRequest, resp *fuse.ReadResponse) error {
	if f.read == nil {
		return syscall.EACCES
	}
	f.RLock()
	defer f.RUnlock()
	res := f.read(f.h)
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
	wrote, err := f.write(f.h, req.Data)
	resp.Size = wrote
	return err
}

// Size return a uint64 measure of the filo size in bytes.
func (f *filo) Size() uint64 {
	if f.read != nil {
		res := f.read(f.h)
		return uint64(res.Len())
	}
	return 0
}
