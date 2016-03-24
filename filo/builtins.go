package filo

var NilFilo Filo = New("", nil, nil)

var builtIns []Filo = []Filo{
	NilFilo,
	New("root_width", rootWidthRead, nil),
	New("root_height", rootHeightRead, nil),
	New("border_width", borderWidthRead, nil),
	New("border_color", nil, borderColorWrite),
	New("area", areaRead, nil),
	New("width", widthRead, widthWrite),
	New("height", heightRead, heightWrite),
	New("X", xRead, xWrite),
	New("Y", yRead, yWrite),
	New("instance", instanceRead, nil),
	New("class", classRead, nil),
	New("mapstate", mapStateRead, nil),
	New("mapped", mappedRead, nil),
	New("viewable", viewableRead, nil),
}
