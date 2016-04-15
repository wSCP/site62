package windows

import "github.com/wSCP/site62/node"

func mkTree() node.Tree {
	return node.NewTree(
		nil,
		node.BaseDir("root", 0700,
			node.BaseDir("geometry", 0700,
				node.BaseFile("width", 0400, "root_width"),
				node.BaseFile("height", 0400, "root_height"),
			)),
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
}

func Block() node.Node {
	t := mkTree()
	return node.New(
		node.Directory,
		"",
		"windows",
		0700,
		nil,
		node.NotAliased,
		t.Head(),
		t.Tail()...,
	)
}
