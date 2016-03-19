package main

import (
	"log"
	"os"
	"path/filepath"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/thrisp/wSCP/xandle"
)

func mount(point string) error {
	x, err := xandle.New("")
	if err != nil {
		return err
	}

	path, err := filepath.Abs(point)
	if err != nil {
		return err
	}

	fileSystem := New(x, path)

	c, err := fuse.Mount(
		path,
	)

	if err != nil {
		return err
	}
	defer c.Close()

	if err := fs.Serve(c, fileSystem); err != nil {
		return err
	}

	<-c.Ready
	if err := c.MountError; err != nil {
		return err
	}

	return nil
}

const (
	success = iota
)

var logger = log.New(os.Stderr, "site62: ", log.Lmicroseconds|log.Llongfile)

func main() {
	err := mount(c.mountPoint)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(success)
}
