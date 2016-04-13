package windows

import (
	"bytes"
	"errors"
	"strconv"

	"github.com/wSCP/site62/filo"
	"github.com/wSCP/site62/state"
	"github.com/wSCP/xandle/x/icccm"
)

func init() {
	filo.Set(windowsFilos...)
}

var windowsFilos []filo.Filo = []filo.Filo{
	filo.New("root_width", rootWidthRead, nil),
	filo.New("root_height", rootHeightRead, nil),
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

func rootWidthRead(s state.State) *bytes.Buffer {
	m, err := getManager(s)
	if err != nil {
		return stringToBuffer(err.Error())
	}
	return stringToBuffer(toStringFnUint32(m.RootWindow().GetWidth))
}

func rootHeightRead(s state.State) *bytes.Buffer {
	m, err := getManager(s)
	if err != nil {
		return stringToBuffer(err.Error())
	}
	return stringToBuffer(toStringFnUint32(m.RootWindow().GetHeight))
}

func borderWidthRead(s state.State) *bytes.Buffer {
	w, err := getWindow(s)
	if err != nil {
		return stringToBuffer(err.Error())
	}
	return stringToBuffer(toStringFnUint32(w.GetBorderWidth))
}

func areaRead(s state.State) *bytes.Buffer {
	w, err := getWindow(s)
	if err != nil {
		return stringToBuffer(err.Error())
	}
	return stringToBuffer(toStringFnUint32(w.GetArea))
}

func widthRead(s state.State) *bytes.Buffer {
	w, err := getWindow(s)
	if err != nil {
		return stringToBuffer(err.Error())
	}
	return stringToBuffer(toStringFnUint32(w.GetWidth))
}

func heightRead(s state.State) *bytes.Buffer {
	w, err := getWindow(s)
	if err != nil {
		return stringToBuffer(err.Error())
	}
	return stringToBuffer(toStringFnUint32(w.GetHeight))
}

func xRead(s state.State) *bytes.Buffer {
	w, err := getWindow(s)
	if err != nil {
		return stringToBuffer(err.Error())
	}
	return stringToBuffer(toStringFnUint32(w.GetX))
}

func yRead(s state.State) *bytes.Buffer {
	w, err := getWindow(s)
	if err != nil {
		return stringToBuffer(err.Error())
	}
	return stringToBuffer(toStringFnUint32(w.GetY))
}

var WmClassGetError = errors.New("unable to get WmClass")

func getWmClass(s state.State) (*icccm.WmClass, error) {
	if w, err := getWindow(s); err == nil {
		res, err := s.WmClassGet(w.Id())
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	return nil, WmClassGetError
}

func instanceRead(s state.State) *bytes.Buffer {
	i, err := getWmClass(s)
	if err != nil {
		return stringToBuffer(err.Error())
	}
	return stringToBuffer(i.Instance)
}

func classRead(s state.State) *bytes.Buffer {
	c, err := getWmClass(s)
	if err != nil {
		return stringToBuffer(err.Error())
	}
	return stringToBuffer(c.Class)
}

func mapStateRead(s state.State) *bytes.Buffer {
	w, err := getWindow(s)
	if err != nil {
		return stringToBuffer(err.Error())
	}
	return stringToBuffer(toStringFnInt(w.MapState))
}

func mappedRead(s state.State) *bytes.Buffer {
	w, err := getWindow(s)
	if err != nil {
		return stringToBuffer(err.Error())
	}
	return stringToBuffer(toStringFnBool(w.Mapped))
}

func viewableRead(s state.State) *bytes.Buffer {
	w, err := getWindow(s)
	if err != nil {
		return stringToBuffer(err.Error())
	}
	return stringToBuffer(toStringFnBool(w.Viewable))
}

func borderColorWrite(s state.State, in []byte) (int, error) {
	return 0, nil
}

func borderWidthWrite(s state.State, in []byte) (int, error) {
	return 0, nil
}

func widthWrite(s state.State, in []byte) (int, error) {
	return 0, nil
}

func heightWrite(s state.State, in []byte) (int, error) {
	return 0, nil
}

func xWrite(s state.State, in []byte) (int, error) {
	return 0, nil
}

func yWrite(s state.State, in []byte) (int, error) {
	return 0, nil
}
