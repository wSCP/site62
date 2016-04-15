package extension

import (
	"reflect"
)

type Inserter interface {
	Insert(...interface{})
}

type Extender interface {
	Extend(...Extension)
	ExtendFunctions(...Function)
}

type Extension interface {
	Tag() string
	Extender
	Functions
	Inserter
	Runner
}

func New(tag string, fns ...Function) Extension {
	e := &extension{
		tag:       tag,
		functions: &functions{},
		holds:     make(map[string]reflect.Value),
	}
	e.ExtendFunctions(fns...)
	e.runner = newRunner(e)
	return e
}

type extension struct {
	tag     string
	inserts []interface{}
	holds   map[string]reflect.Value
	*functions
	*runner
}

func (e *extension) Tag() string {
	return e.tag
}

func (e *extension) Extend(extensions ...Extension) {
	for _, extension := range extensions {
		e.ExtendFunctions(extension.Functions()...)
	}
}

func goodFunc(typ reflect.Type) bool {
	switch {
	case typ.NumOut() == 1:
		return true
	case typ.NumOut() == 2 && typ.Out(1) == rferrorType:
		return true
	}
	return false
}

func valueFunc(fn interface{}) reflect.Value {
	v := reflect.ValueOf(fn)
	if v.Kind() != reflect.Func {
		panic(NotAFunction(fn, fn))
	}
	if !goodFunc(v.Type()) {
		panic(BadFunc(fn, v.Type().NumOut()))
	}
	return v
}

func (e *extension) ExtendFunctions(fns ...Function) {
	for _, fn := range fns {
		e.holds[fn.Key()] = valueFunc(fn.Value())
		e.Add(fn)
	}
}

func (e *extension) Insert(args ...interface{}) {
	e.inserts = args
}
