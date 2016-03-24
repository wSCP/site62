package node

import (
	"strconv"

	"bazil.org/fuse"
	"github.com/thrisp/wSCP/xandle"
	"github.com/thrisp/wSCP/xandle/monitor"
	"github.com/thrisp/wSCP/xandle/window"
)

type AliasFindFn func(string, Node, Node) (bool, Node)

type AliasListFn func(Node, xandle.Xandle) []fuse.Dirent

type Aliaser interface {
	AliasKey() string
	Aliased() bool
	Find(string, Node, Node) (bool, Node)
	List(Node, xandle.Xandle) []fuse.Dirent
}

type aliaser struct {
	key     string
	aliased bool
	find    AliasFindFn
	list    AliasListFn
}

func NewAliaser(k string, a bool, f AliasFindFn, l AliasListFn) Aliaser {
	return &aliaser{k, a, f, l}
}

func (a *aliaser) AliasKey() string {
	return a.key
}

func (a *aliaser) Aliased() bool {
	return a.aliased
}

func (a *aliaser) Find(requested string, h Node, n Node) (bool, Node) {
	return a.find(requested, h, n)
}

func (a *aliaser) List(n Node, x xandle.Xandle) []fuse.Dirent {
	return a.list(n, x)
}

type aliasers map[string]Aliaser

func (as aliasers) SetAliasers(a ...Aliaser) {
	for _, al := range a {
		as[al.AliasKey()] = al
	}
}

func (as aliasers) GetAliaser(key string) Aliaser {
	if ret, exists := as[key]; exists {
		return ret
	}
	return NotAliased
}

var Aliasers aliasers

func init() {
	Aliasers = make(aliasers)
	Aliasers.SetAliasers(NotAliased, WindowAliaser, MonitorAliaser)
}

var (
	NotAliased     Aliaser = NewAliaser("not_aliased", false, nil, nil)
	WindowAliaser  Aliaser = NewAliaser("window_aliaser", true, findWindowDir, listWindowDir)
	MonitorAliaser Aliaser = NewAliaser("monitor_aliaser", true, findMonitorDir, listMonitorDir)
)

func findWindowDir(requested string, head Node, n Node) (bool, Node) {
	x := head.Xandle()
	if wid, exists := x.WindowExists(requested); exists {
		rd := n.Copy()
		rd.SetName(requested)
		rd.SetHead(head)
		rd.SetXandle(x)
		rd.SetWindow(window.New(x.Conn(), wid, x.Root()))
		return true, rd
	}
	return false, nil
}

func listWindowDir(n Node, x xandle.Xandle) []fuse.Dirent {
	var ret []fuse.Dirent
	wl, err := x.ListWindows()
	if err == nil {
		for _, w := range wl {
			ret = append(ret, n.Entry(windowString(w)))
		}
	}
	return ret
}

func getMonitors(n Node) *monitor.Monitors {
	x := n.Xandle()
	m := x.Monitors()
	return m
}

func findMonitorDir(requested string, head Node, n Node) (bool, Node) {
	if id, err := strconv.ParseUint(requested, 10, 32); err == nil {
		if m, exists := monitor.Exists(uint32(id), getMonitors(head)); exists {
			rd := n.Copy()
			rd.SetName(requested)
			rd.SetHead(head)
			rd.SetXandle(head.Xandle())
			rd.SetMonitor(m)
			return true, rd
		}
	}
	return false, nil
}

func listMonitorDir(n Node, x xandle.Xandle) []fuse.Dirent {
	var ret []fuse.Dirent
	ms := getMonitors(n)
	for _, m := range monitor.All(ms) {
		ret = append(ret, n.Entry(strconv.FormatUint(uint64(m.Id()), 10)))
	}
	return ret
}
