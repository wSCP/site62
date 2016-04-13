package filesystem

import (
	"log"
	"os"
	"sort"

	we "github.com/wSCP/site62/extensions/windows"
	"github.com/wSCP/site62/node"
	"github.com/wSCP/site62/state"
	"github.com/wSCP/xandle"
	"github.com/wSCP/xandle/x"
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
	config{1001, DefaultState},
	config{1002, DefaultRootFn},
	config{1003, DefaultLogger},
	config{1004, windows},
}

func DefaultMountPoint(f *FS) error {
	if f.RootPath == "" {
		f.RootPath = "/tmp/site62"
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

func defaultX() x.Handle {
	//es := filepath.Join(f.RootPath, "event")
	h, err := x.New("")
	if err != nil {
		panic(err)
	}
	return h
}

func DefaultState(f *FS) error {
	if f.state == nil {
		f.state = state.DefaultStateFn(xandle.New(defaultX()))
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

func DefaultLogger(f *FS) error {
	if f.Logger == nil {
		f.Logger = log.New(os.Stderr, "site62:fs: ", log.Lmicroseconds|log.Llongfile)
	}
	return nil
}

func Logger(l *log.Logger) Config {
	return config{
		50,
		func(f *FS) error {
			f.Logger = l
			return nil
		},
	}
}

func windows(f *FS) error {
	f.NewMounts("windows", "/", we.Block.Copy())
	s := f.state
	s.Extend(we.New(s.Conn(), s.Root()))
	return nil
}
