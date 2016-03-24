package filo

import (
	"bytes"
	"errors"
	"strconv"

	"github.com/thrisp/wSCP/xandle"
	"github.com/thrisp/wSCP/xandle/window"
)

// A function used internally by a File to read.
type FileRead func(Header) *bytes.Buffer

func stringToBuffer(s string) *bytes.Buffer {
	var b = new(bytes.Buffer)
	b.WriteString(s)
	return b
}

func toStringFnInt(fn func() int) string {
	return strconv.FormatInt(int64(fn()), 10)
}

func toStringFnUint32(fn func() uint32) string {
	return strconv.FormatInt(int64(fn()), 10)
}

func toStringFnBool(fn func() bool) string {
	return strconv.FormatBool(fn())
}

var NoXandle = errors.New("No xandle available")

func getXandle(h Header) (xandle.Xandle, error) {
	if x := h.Xandle(); x != nil {
		return x, nil
	}
	return nil, NoXandle
}

var NoWindow = errors.New("No window available")

func getWindow(h Header) (window.Window, error) {
	if w := h.Window(); w != nil {
		return w, nil
	}
	return nil, NoWindow
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

func borderWidthRead(h Header) *bytes.Buffer {
	if w, err := getWindow(h); err == nil {
		return stringToBuffer(toStringFnUint32(w.GetBorderWidth))
	}
	return nil
}

func areaRead(h Header) *bytes.Buffer {
	if w, err := getWindow(h); err == nil {
		return stringToBuffer(toStringFnUint32(w.GetArea))
	}
	return nil
}

func widthRead(h Header) *bytes.Buffer {
	if w, err := getWindow(h); err == nil {
		return stringToBuffer(toStringFnUint32(w.GetWidth))
	}
	return nil
}

func heightRead(h Header) *bytes.Buffer {
	if w, err := getWindow(h); err == nil {
		return stringToBuffer(toStringFnUint32(w.GetHeight))
	}
	return nil
}

func xRead(h Header) *bytes.Buffer {
	if w, err := getWindow(h); err == nil {
		return stringToBuffer(toStringFnUint32(w.GetX))
	}
	return nil
}

func yRead(h Header) *bytes.Buffer {
	if w, err := getWindow(h); err == nil {
		return stringToBuffer(toStringFnUint32(w.GetY))
	}
	return nil
}

func instanceRead(h Header) *bytes.Buffer {
	if x, err := getXandle(h); err == nil {
		if w, err := getWindow(h); err == nil {
			res, err := x.WmClassGet(w.Id())
			if err != nil {
				return stringToBuffer(err.Error())
			}
			return stringToBuffer(res.Instance)
		}
	}
	return nil
}

func classRead(h Header) *bytes.Buffer {
	if x, err := getXandle(h); err == nil {
		if w, err := getWindow(h); err == nil {
			res, err := x.WmClassGet(w.Id())
			if err != nil {
				return stringToBuffer(err.Error())
			}
			return stringToBuffer(res.Class)
		}
	}
	return nil
}

func mapStateRead(h Header) *bytes.Buffer {
	if w, err := getWindow(h); err == nil {
		return stringToBuffer(toStringFnInt(w.MapState))
	}
	return nil
}

func mappedRead(h Header) *bytes.Buffer {
	if w, err := getWindow(h); err == nil {
		return stringToBuffer(toStringFnBool(w.Mapped))
	}
	return nil
}

func viewableRead(h Header) *bytes.Buffer {
	if w, err := getWindow(h); err == nil {
		return stringToBuffer(toStringFnBool(w.Viewable))
	}
	return nil
}
