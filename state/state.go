package state

import (
	"github.com/wSCP/site62/extension"
	"github.com/wSCP/xandle"
)

type StateFn func(xandle.Xandle) State

func DefaultStateFn(x xandle.Xandle) State {
	return &state{
		Xandle:    x,
		Extension: extension.New("state-extension"),
	}
}

type Identifier interface {
	Identity() string
	SetIdentity(string)
}

type State interface {
	Identifier
	xandle.Xandle
	extension.Extension
}

type state struct {
	id string
	xandle.Xandle
	extension.Extension
}

func (s *state) Identity() string {
	return s.id
}

func (s *state) SetIdentity(id string) {
	s.id = id
}
