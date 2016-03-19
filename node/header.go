package node

import (
	"github.com/thrisp/wSCP/xandle"
	"github.com/thrisp/wSCP/xandle/window"
)

type Header interface {
	Xandle() xandle.Xandle
	SetXandle(xandle.Xandle)
	Window() window.Window
	SetWindow(window.Window)
}

type header struct {
	x xandle.Xandle
	w window.Window
}

func NewHeader(x xandle.Xandle, w window.Window) Header {
	return &header{x, w}
}

func (h *header) Xandle() xandle.Xandle {
	return h.x
}

func (h *header) SetXandle(x xandle.Xandle) {
	h.x = x
}

func (h *header) Window() window.Window {
	return h.w
}

func (h *header) SetWindow(w window.Window) {
	h.w = w
}
