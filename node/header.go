package node

import (
	"github.com/wSCP/xandle"
	"github.com/wSCP/xandle/monitor"
	"github.com/wSCP/xandle/tag"
	"github.com/wSCP/xandle/window"
)

type Header interface {
	CurrentHeader() Header
	Xandle() xandle.Xandle
	SetXandle(xandle.Xandle)
	Window() window.Window
	SetWindow(window.Window)
	Monitor() monitor.Monitor
	SetMonitor(monitor.Monitor)
	Tag() tag.Tag
	SetTag(tag.Tag)
}

type header struct {
	x xandle.Xandle
	w window.Window
	m monitor.Monitor
	t tag.Tag
}

var EmptyHeader header = header{nil, nil, nil, nil}

func NewHeader(x xandle.Xandle, w window.Window, m monitor.Monitor, t tag.Tag) Header {
	return &header{x, w, m, t}
}

func (h *header) CurrentHeader() Header {
	return h
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

func (h *header) Monitor() monitor.Monitor {
	return h.m
}

func (h *header) SetMonitor(m monitor.Monitor) {
	h.m = m
}

func (h *header) Tag() tag.Tag {
	return h.t
}

func (h *header) SetTag(t tag.Tag) {
	h.t = t
}
