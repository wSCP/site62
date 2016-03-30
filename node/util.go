package node

import "bazil.org/fuse"

func entry(name string, t fuse.DirentType) fuse.Dirent {
	return fuse.Dirent{
		Type: t,
		Name: name,
	}
}
