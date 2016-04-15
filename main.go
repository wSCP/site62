package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/wSCP/site62/filesystem"
)

var c *configuration

type configuration struct {
	mountPoint string
	fns        []filesystem.Config
}

func (c *configuration) parse() {
	path, err := filepath.Abs(c.mountPoint)
	if err != nil {
		l.Fatal(err)
	}
	c.fns = append(c.fns, filesystem.Logger(l), filesystem.MountPoint(path))
}

func init() {
	c = &configuration{}
	flag.StringVar(&c.mountPoint, "mount", "/tmp/site62", "A string path to mount the file system to.")
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
