package extension

import "reflect"

type Extension interface {
	Tag() string
	Insert(...interface{})
	Run(string, ...interface{}) (interface{}, error)
	MustRun(string, ...interface{}) interface{}
	SetExtensionFunctions(...Function)
	Extend(...Extension)
	All() []Function
	Returns
}

func New(tag string, fns ...Function) Extension {
	e := &extension{
		tag:        tag,
		extensions: make(map[string]reflect.Value),
	}
	e.SetExtensionFunctions(fns...)
	e.Returns = &returns{e}
	return e
}

type extension struct {
	tag        string
	inserts    []interface{}
	extensions map[string]reflect.Value
	functions  []Function
	Returns
}

func (e *extension) Tag() string {
	return e.tag
}

func (e *extension) Insert(args ...interface{}) {
	e.inserts = args
}

func (e *extension) Run(fnName string, arg ...interface{}) (interface{}, error) {
	if fn, ok := e.extensions[fnName]; ok {
		var args []interface{}
		args = append(args, e.inserts...)
		args = append(args, arg...)
		return call(fn, args...)
	}
	return nil, NotAnExtension(fnName)
}

func (e *extension) MustRun(fnName string, arg ...interface{}) interface{} {
	var ret interface{}
	var err error
	if ret, err = e.Run(fnName, arg...); err != nil {
		panic(err.Error())
	}
	return ret
}

func (e *extension) SetExtensionFunctions(fns ...Function) {
	for _, fn := range fns {
		e.extensions[fn.Key()] = valueFunc(fn.Value())
		e.functions = append(e.functions, fn)
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

func (e *extension) Extend(extensions ...Extension) {
	for _, extension := range extensions {
		e.SetExtensionFunctions(extension.All()...)
	}
}

func (e *extension) All() []Function {
	return e.functions
}
