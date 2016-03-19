package node

import (
	"strconv"

	"bazil.org/fuse"

	"github.com/BurntSushi/xgb/xproto"
)

func windowString(w xproto.Window) string {
	return strconv.FormatUint(uint64(w), 10)
}

func entry(name string, t fuse.DirentType) fuse.Dirent {
	return fuse.Dirent{
		Type: t,
		Name: name,
	}
}
