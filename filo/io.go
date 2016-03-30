package filo

import (
	"bytes"
	"errors"
	"strconv"

	"github.com/wSCP/xandle"
)

// A function used internally by a File to read.
type FileRead func(Header) *bytes.Buffer

// A function used internally by a File to write.
type FileWrite func(Header, []byte) (int, error)

var NilFilo Filo = New("", nil, nil)

var builtIns []Filo = []Filo{
	NilFilo,
	New("root_width", rootWidthRead, nil),
	New("root_height", rootHeightRead, nil),
}

func stringToBuffer(s string) *bytes.Buffer {
	var b = new(bytes.Buffer)
	b.WriteString(s)
	return b
}

func toStringFnUint32(fn func() uint32) string {
	return strconv.FormatInt(int64(fn()), 10)
}

var NoXandle = errors.New("No xandle available")

func getXandle(h Header) (xandle.Xandle, error) {
	if x := h.Xandle(); x != nil {
		return x, nil
	}
	return nil, NoXandle
}

func rootWidthRead(h Header) *bytes.Buffer {
	if x, err := getXandle(h); err == nil {
		return stringToBuffer(toStringFnUint32(x.RootWindow().GetWidth))
	}
	return nil
}

func rootHeightRead(h Header) *bytes.Buffer {
	if x, err := getXandle(h); err == nil {
		return stringToBuffer(toStringFnUint32(x.RootWindow().GetHeight))
	}
	return nil
}
