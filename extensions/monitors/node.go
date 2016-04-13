package monitors

/*
func Monitors(f *FS) error {
	f.NewMounts(monitors, &node.MonitorsBlock)
	return nil
}

var MonitorsBlock node = node{
	Header:  &EmptyHeader,
	Kind:    kind{Directory},
	name:    "monitors",
	mode:    os.ModeDir | 0700,
	Aliaser: NotAliased,
	Tree: NewTree(nil, baseAliasedDir("0x*", 0700, MonitorAliaser,
		baseDir("geometry", 0700,
			baseFile("area", 0600, ""),
			baseFile("width", 0600, ""),
			baseFile("height", 0600, ""),
			baseFile("x", 0600, ""),
			baseFile("y", 0600, ""),
		),
		baseFile("root", 0400, ""),
		baseFile("wired", 0400, ""),
		baseFile("focus", 0400, ""),
	),
		baseFile("focus", 0600, ""),
		baseFile("primary", 0600, ""),
	),
}

MonitorAliaser Aliaser = NewAliaser("monitor_aliaser", true, findMonitorDir, listMonitorDir)

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
			//rd.SetTail(TagsBlock(rd.CurrentHeader()))
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
*/
