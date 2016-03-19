package filo

var builtIns []Filo = []Filo{
	New("", nil, nil),
	New("root_width", rootWidthRead, nil),
	New("root_height", rootHeightRead, nil),
	New("border_width", borderWidthRead, nil),
	New("border_color", nil, borderColorWrite),
	New("width", widthRead, widthWrite),
	New("height", heightRead, heightWrite),
	New("X", xRead, xWrite),
	New("Y", yRead, yWrite),
	New("instance", instanceRead, nil),
	New("class", classRead, nil),
}
