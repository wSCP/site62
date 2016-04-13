package extension

import "fmt"

type extensionError struct {
	err  string
	vals []interface{}
}

func (m *extensionError) Error() string {
	return fmt.Sprintf("%s", fmt.Sprintf(m.err, m.vals...))
}

func (m *extensionError) Out(vals ...interface{}) *extensionError {
	m.vals = vals
	return m
}

func Xrror(err string) *extensionError {
	return &extensionError{err: err}
}
