package filo

import (
	"github.com/thrisp/wSCP/xandle"
	"github.com/thrisp/wSCP/xandle/window"
)

// A function used internally by a File to write.
type FileWrite func(window.Window, xandle.Xandle, []byte) (int, error)

func borderColorWrite(w window.Window, x xandle.Xandle, in []byte) (int, error) {
	return 0, nil
}

func borderWidthWrite(w window.Window, x xandle.Xandle, in []byte) (int, error) {
	return 0, nil
}

func widthWrite(w window.Window, x xandle.Xandle, in []byte) (int, error) {
	return 0, nil
}

func heightWrite(w window.Window, x xandle.Xandle, in []byte) (int, error) {
	return 0, nil
}

func xWrite(w window.Window, x xandle.Xandle, in []byte) (int, error) {
	return 0, nil
}

func yWrite(w window.Window, x xandle.Xandle, in []byte) (int, error) {
	return 0, nil
}
