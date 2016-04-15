package windows

import (
	"errors"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	w "github.com/wSCP/arboreal/window"
	"github.com/wSCP/site62/extension"
	"github.com/wSCP/site62/state"
)

type Manager struct {
	w.Windows
	m map[string]w.Window
}

func New(c *xgb.Conn, r xproto.Window) extension.Extension {
	m := &Manager{
		Windows: w.NewWindows(c, r),
		m:       make(map[string]w.Window),
	}
	get := extension.NewFunction("windows-manager", func() *Manager { return m })
	return extension.New("WINDOWS", get)
}

var NoManager error = errors.New("no manager.")

func getManager(s state.State) (*Manager, error) {
	if mr, err := s.Run("windows-manager"); err == nil {
		if m, ok := mr.(*Manager); ok {
			return m, nil
		}
	}
	return nil, NoManager
}

var NoWindow error = errors.New("no window.")

func getWindow(s state.State) (w.Window, error) {
	if m, err := getManager(s); err == nil {
		return m.getWindow(s.Identity())
	}
	return nil, NoWindow
}

func (m *Manager) getWindow(id string) (w.Window, error) {
	if w, exists := m.m[id]; exists {
		return w, nil
	}
	return nil, NoWindow
}
