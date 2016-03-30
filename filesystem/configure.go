package filesystem

import (
	"path/filepath"
	"sort"

	ws "github.com/wSCP/site62/modules/windows"
	"github.com/wSCP/site62/node"
	"github.com/wSCP/xandle"
)

type ConfigFn func(*FS) error

type Config interface {
	Order() int
	Configure(*FS) error
}

type config struct {
	order int
	fn    ConfigFn
}

func DefaultConfig(fn ConfigFn) Config {
	return config{50, fn}
}

func NewConfig(order int, fn ConfigFn) Config {
	return config{order, fn}
}

func (c config) Order() int {
	return c.order
}

func (c config) Configure(f *FS) error {
	return c.fn(f)
}

type configList []Config

func (c configList) Len() int {
	return len(c)
}

func (c configList) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c configList) Less(i, j int) bool {
	return c[i].Order() < c[j].Order()
}

type Configuration interface {
	Add(...Config)
	AddFn(...ConfigFn)
	Configure() error
	Configured() bool
}

type configuration struct {
	f          *FS
	configured bool
	list       configList
}

func newConfiguration(f *FS, conf ...Config) *configuration {
	c := &configuration{
		f:    f,
		list: builtIns,
	}
	c.Add(conf...)
	return c
}

func (c *configuration) Add(conf ...Config) {
	c.list = append(c.list, conf...)
}

func (c *configuration) AddFn(fns ...ConfigFn) {
	for _, fn := range fns {
		c.list = append(c.list, DefaultConfig(fn))
	}
}

func configure(f *FS, conf ...Config) error {
	for _, c := range conf {
		err := c.Configure(f)
		if err != nil {
			return err
		}
	}
	return nil
}

func respondTo(c *configuration, err error) {
	if err != nil {
		c.f.Fatalf("%s", err.Error())
	}
}

func (c *configuration) Configure() error {
	sort.Sort(c.list)

	err := configure(c.f, c.list...)
	respondTo(c, err)
	if err == nil {
		c.configured = true
	}

	return err
}

func (c *configuration) Configured() bool {
	return c.configured
}

var builtIns = []Config{
	config{1000, DefaultMountPoint},
	config{1001, DefaultXandle},
	config{1002, DefaultRootFn},
	config{50, windows},
}

func DefaultMountPoint(f *FS) error {
	if f.RootPath == "" {
		f.RootPath = "/tmp/site62"
	}
	return nil
}

func DefaultXandle(f *FS) error {
	if f.Xandle == nil {
		es := filepath.Join(f.RootPath, "event")
		x, err := xandle.New("", es)
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

func MountPoint(path string) Config {
	return config{
		50,
		func(f *FS) error {
			f.RootPath = path
			return nil
		},
	}
}

func windows(f *FS) error {
	f.NewMounts("windows", "/", ws.WindowsBlock.Copy())
	return nil
}
