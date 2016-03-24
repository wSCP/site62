package filesystem

import (
	"github.com/thrisp/wSCP/site62/node"
	"github.com/thrisp/wSCP/xandle"
)

func (f *FS) Configure() error {
	var err error
	cfs := [][]ConfigureFn{f.fns, f.dfns}
	for _, cf := range cfs {
		for _, fn := range cf {
			err = fn(f)
			if err != nil {
				return err
			}
		}
	}
	f.configured = true
	return nil
}

type Configuration struct {
	configured bool
	fns        []ConfigureFn
	dfns       []ConfigureFn
	verbose    bool
	RootFn     func(xandle.Xandle, string) node.Node
	RootPath   string
	mounts     []Mounts
}

func NewConfiguration(c ...ConfigureFn) *Configuration {
	return &Configuration{
		fns:    c,
		dfns:   []ConfigureFn{DefaultXandle, DefaultRootFn},
		mounts: make([]Mounts, 0),
	}
}

type ConfigureFn func(*FS) error

func DefaultXandle(f *FS) error {
	if f.Xandle == nil {
		x, err := xandle.New("")
		if err != nil {
			return err
		}
		f.Xandle = x
		return nil
	}
	return nil
}

func DefaultRootFn(f *FS) error {
	if f.RootFn == nil {
		f.RootFn = node.NewRoot
	}
	return nil
}

func MountPoint(path string) ConfigureFn {
	return func(f *FS) error {
		f.RootPath = path
		return nil
	}
}

func Clients(f *FS) error {
	f.NewMounts(clients, &node.ClientsBlock)
	return nil
}

func Monitors(f *FS) error {
	f.NewMounts(monitors, &node.MonitorsBlock)
	return nil
}
