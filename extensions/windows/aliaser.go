package windows

import (
	"strconv"

	"bazil.org/fuse"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/wSCP/arboreal/window"
	"github.com/wSCP/site62/node"
	"github.com/wSCP/site62/state"
)

var WindowAliaser node.Aliaser = node.NewAliaser("window_aliaser", true, findWindowDir, listWindowDir)

func findWindowDir(requested string, head node.Node, n node.Node) (bool, node.Node) {
	s := head.Current()
	if m, err := getManager(s); err == nil {
		if wid, exists := m.WindowExists(requested); exists {
			rd := n.Copy()
			rd.SetName(requested)
			rd.SetState(s)
			rds := rd.Current()
			rds.SetIdentity(requested)
			m.m[requested] = window.New(s.Conn(), wid, s.Root())
			return true, rd
		}
	}
	return false, nil
}

func windowString(w xproto.Window) string {
	return strconv.FormatUint(uint64(w), 10)
}

func listWindowDir(n node.Node, s state.State) []fuse.Dirent {
	var ret []fuse.Dirent
	if m, err := getManager(s); err == nil {
		wl, err := m.ListWindows()
		if err == nil {
			for _, w := range wl {
				ret = append(ret, n.Entry(windowString(w)))
			}
		}
	}
	return ret
}
