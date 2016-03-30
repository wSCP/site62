package windows

import (
	"strconv"

	"bazil.org/fuse"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/wSCP/site62/node"
	"github.com/wSCP/xandle"
	w "github.com/wSCP/xandle/window"
)

var tree node.Tree = node.NewTree(
	nil,
	node.BaseAliasedDir("x*", 0700, WindowAliaser,
		node.BaseDir("border", 0700,
			node.BaseFile("color", 0200, "border_color"),
			node.BaseFile("width", 0600, "border_width"),
		),
		node.BaseDir("geometry", 0700,
			node.BaseFile("area", 0600, "area"),
			node.BaseFile("width", 0600, "width"),
			node.BaseFile("height", 0600, "height"),
			node.BaseFile("x", 0600, "X"),
			node.BaseFile("y", 0600, "Y"),
		),
		node.BaseDir("state", 0700,
			node.BaseFile("map_state", 0600, "mapstate"),
			node.BaseFile("mapped", 0600, "mapped"),
			node.BaseFile("focus", 0600, ""),
			node.BaseFile("ignore", 0600, ""),
			node.BaseFile("stack", 0200, ""),
		),
		node.BaseFile("instance", 0400, "instance"),
		node.BaseFile("class", 0400, "class"),
	),
)

var WindowsBlock = node.New(
	node.Directory,
	"",
	"windows",
	0700,
	nil,
	node.NotAliased,
	tree.Head(),
	tree.Tail()...,
)

var WindowAliaser node.Aliaser = node.NewAliaser("window_aliaser", true, findWindowDir, listWindowDir)

func findWindowDir(requested string, head node.Node, n node.Node) (bool, node.Node) {
	x := head.Xandle()
	if wid, exists := x.WindowExists(requested); exists {
		rd := n.Copy()
		rd.SetName(requested)
		rd.SetHead(head)
		rd.SetXandle(x)
		rd.SetWindow(w.New(x.Conn(), wid, x.Root()))
		return true, rd
	}
	return false, nil
}

func windowString(w xproto.Window) string {
	return strconv.FormatUint(uint64(w), 10)
}

func listWindowDir(n node.Node, x xandle.Xandle) []fuse.Dirent {
	var ret []fuse.Dirent
	wl, err := x.ListWindows()
	if err == nil {
		for _, w := range wl {
			ret = append(ret, n.Entry(windowString(w)))
		}
	}
	return ret
}
