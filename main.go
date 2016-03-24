package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/thrisp/wSCP/site62/filesystem"
)

var c *configuration

type configuration struct {
	mountPoint string
	monitors   bool
	tags       bool
	clients    bool
	fns        []filesystem.ConfigureFn
}

func (c *configuration) parse() {
	path, err := filepath.Abs(c.mountPoint)
	if err != nil {
		l.Fatal(err)
	}
	c.fns = append(c.fns, filesystem.MountPoint(path))
	if c.monitors {
		c.fns = append(c.fns, filesystem.Monitors)
	}
	if c.tags {
		//c.fns = append(c.fns, filesystem.Tags)
	}
	if c.clients {
		c.fns = append(c.fns, filesystem.Clients)
	}
}

func init() {
	c = &configuration{}
	flag.StringVar(&c.mountPoint, "mount", "/tmp/site62", "A string path to mount the file system to.")
	flag.BoolVar(&c.monitors, "monitors", false, "Specify the file system to mount access to monitors")
	flag.BoolVar(&c.tags, "tags", false, "Specify the file system to mount access to tags")
	flag.BoolVar(&c.clients, "clients", false, "Specify the file system to mount access to clients")
	flag.Parse()
	c.parse()
}

var l = log.New(os.Stderr, "site62: ", log.Lmicroseconds|log.Llongfile)

func main() {
	var err error

	fis := filesystem.New(c.fns...)

	err = fis.Mount()
	if err != nil {
		l.Fatal(err)
	}

	os.Exit(0)
}
