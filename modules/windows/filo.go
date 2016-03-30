package windows

import (
	"bytes"
	"errors"
	"strconv"

	"github.com/wSCP/site62/filo"
	"github.com/wSCP/xandle"
	"github.com/wSCP/xandle/window"
)

func init() {
	filo.Set(windowsFilos...)
}

var windowsFilos []filo.Filo = []filo.Filo{
	filo.New("border_width", borderWidthRead, nil),
	filo.New("border_color", nil, borderColorWrite),
	filo.New("area", areaRead, nil),
	filo.New("width", widthRead, widthWrite),
	filo.New("height", heightRead, heightWrite),
	filo.New("X", xRead, xWrite),
	filo.New("Y", yRead, yWrite),
	filo.New("instance", instanceRead, nil),
	filo.New("class", classRead, nil),
	filo.New("mapstate", mapStateRead, nil),
	filo.New("mapped", mappedRead, nil),
	filo.New("viewable", viewableRead, nil),
}

var NoXandle = errors.New("No xandle available")

func getXandle(h filo.Header) (xandle.Xandle, error) {
	if x := h.Xandle(); x != nil {
		return x, nil
	}
	return nil, NoXandle
}

var NoWindow = errors.New("No window available")

func getWindow(h filo.Header) (window.Window, error) {
	if w := h.Window(); w != nil {
		return w, nil
	}
	return nil, NoWindow
}

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

func borderWidthRead(h filo.Header) *bytes.Buffer {
	if w, err := getWindow(h); err == nil {
		return stringToBuffer(toStringFnUint32(w.GetBorderWidth))
	}
	return nil
}

func areaRead(h filo.Header) *bytes.Buffer {
	if w, err := getWindow(h); err == nil {
		return stringToBuffer(toStringFnUint32(w.GetArea))
	}
	return nil
}

func widthRead(h filo.Header) *bytes.Buffer {
	if w, err := getWindow(h); err == nil {
		return stringToBuffer(toStringFnUint32(w.GetWidth))
	}
	return nil
}

func heightRead(h filo.Header) *bytes.Buffer {
	if w, err := getWindow(h); err == nil {
		return stringToBuffer(toStringFnUint32(w.GetHeight))
	}
	return nil
}

func xRead(h filo.Header) *bytes.Buffer {
	if w, err := getWindow(h); err == nil {
		return stringToBuffer(toStringFnUint32(w.GetX))
	}
	return nil
}

func yRead(h filo.Header) *bytes.Buffer {
	if w, err := getWindow(h); err == nil {
		return stringToBuffer(toStringFnUint32(w.GetY))
	}
	return nil
}

func instanceRead(h filo.Header) *bytes.Buffer {
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

func classRead(h filo.Header) *bytes.Buffer {
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

func mapStateRead(h filo.Header) *bytes.Buffer {
	if w, err := getWindow(h); err == nil {
		return stringToBuffer(toStringFnInt(w.MapState))
	}
	return nil
}

func mappedRead(h filo.Header) *bytes.Buffer {
	if w, err := getWindow(h); err == nil {
		return stringToBuffer(toStringFnBool(w.Mapped))
	}
	return nil
}

func viewableRead(h filo.Header) *bytes.Buffer {
	if w, err := getWindow(h); err == nil {
		return stringToBuffer(toStringFnBool(w.Viewable))
	}
	return nil
}

func borderColorWrite(h filo.Header, in []byte) (int, error) {
	return 0, nil
}

func borderWidthWrite(h filo.Header, in []byte) (int, error) {
	return 0, nil
}

func widthWrite(h filo.Header, in []byte) (int, error) {
	return 0, nil
}

func heightWrite(h filo.Header, in []byte) (int, error) {
	return 0, nil
}

func xWrite(h filo.Header, in []byte) (int, error) {
	return 0, nil
}

func yWrite(h filo.Header, in []byte) (int, error) {
	return 0, nil
}
