package filo

import (
	"bytes"
	"strconv"

	"github.com/thrisp/wSCP/xandle"
	"github.com/thrisp/wSCP/xandle/window"
)

// A function used internally by a File to read.
type FileRead func(window.Window, xandle.Xandle) *bytes.Buffer

func toBuffer(s string) *bytes.Buffer {
	var b = new(bytes.Buffer)
	b.WriteString(s)
	return b
}

func toStringFnUint32(fn func() uint32) string {
	return strconv.FormatInt(int64(fn()), 10)
}

func rootWidthRead(w window.Window, x xandle.Xandle) *bytes.Buffer {
	return toBuffer(toStringFnUint32(x.RootWindow().GetWidth))
}

func rootHeightRead(w window.Window, x xandle.Xandle) *bytes.Buffer {
	return toBuffer(toStringFnUint32(x.RootWindow().GetHeight))
}

func borderWidthRead(w window.Window, x xandle.Xandle) *bytes.Buffer {
	return toBuffer(toStringFnUint32(w.GetBorderWidth))
}

func widthRead(w window.Window, x xandle.Xandle) *bytes.Buffer {
	return toBuffer(toStringFnUint32(w.GetWidth))
}

func heightRead(w window.Window, x xandle.Xandle) *bytes.Buffer {
	return toBuffer(toStringFnUint32(w.GetHeight))
}

func xRead(w window.Window, x xandle.Xandle) *bytes.Buffer {
	return toBuffer(toStringFnUint32(w.GetX))
}

func yRead(w window.Window, x xandle.Xandle) *bytes.Buffer {
	return toBuffer(toStringFnUint32(w.GetY))
}

func instanceRead(w window.Window, x xandle.Xandle) *bytes.Buffer {
	res, err := x.WmClassGet(w.Id())
	if err != nil {
		return toBuffer(err.Error())
	}
	return toBuffer(res.Instance)
}

func classRead(w window.Window, x xandle.Xandle) *bytes.Buffer {
	res, err := x.WmClassGet(w.Id())
	if err != nil {
		return toBuffer(err.Error())
	}
	return toBuffer(res.Class)
}
