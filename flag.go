package main

import (
	"flag"

	"github.com/davecgh/go-spew/spew"
)

var c *configuration

type configuration struct {
	mountPoint string
	monitors   bool
	tags       bool
}

func init() {
	c = &configuration{}
	flag.StringVar(&c.mountPoint, "mount", "/tmp/site62", "specify a path to mount the file system to, default is '/tmp/site62'")
	flag.BoolVar(&c.monitors, "monitors", false, "specify access to monitors, default is false")
	flag.BoolVar(&c.tags, "tags", false, "specify access to tags, default is false")
	spew.Dump(c)
	flag.Parse()
}
