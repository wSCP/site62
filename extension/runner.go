package extension

type Runner interface {
	Run(string, ...interface{}) (interface{}, error)
	MustRun(string, ...interface{}) interface{}
	TypeReturner
}

type runner struct {
	e *extension
	*typeReturner
}

func newRunner(e *extension) *runner {
	r := &runner{
		e: e,
	}
	r.typeReturner = &typeReturner{r}
	return r
}

func (r *runner) Run(fnName string, arg ...interface{}) (interface{}, error) {
	if fn, ok := r.e.holds[fnName]; ok {
		var args []interface{}
		args = append(args, r.e.inserts...)
		args = append(args, arg...)
		return call(fn, args...)
	}
	return nil, NotAnExtension(fnName)
}

func (r *runner) MustRun(fnName string, arg ...interface{}) interface{} {
	var ret interface{}
	var err error
	if ret, err = r.e.Run(fnName, arg...); err != nil {
		panic(err.Error())
	}
	return ret
}

type TypeReturner interface {
	RunString(string, ...interface{}) string
	RunInteger(string, ...interface{}) int
	RunBoolean(string, ...interface{}) bool
}

type typeReturner struct {
	r Runner
}

func (t *typeReturner) RunString(name string, arg ...interface{}) string {
	var ret string
	var ok bool
	res := t.r.MustRun(name, arg...)
	if ret, ok = res.(string); !ok {
		panic(NotExpectedReturn(ret, "string").Error())
	}
	return ret
}

func (t *typeReturner) RunInteger(name string, arg ...interface{}) int {
	var ret int
	var ok bool
	res := t.r.MustRun(name, arg...)
	if ret, ok = res.(int); !ok {
		panic(NotExpectedReturn(ret, "integer").Error())
	}
	return ret
}

func (t *typeReturner) RunBoolean(name string, arg ...interface{}) bool {
	var ret bool
	var ok bool
	res := t.r.MustRun(name, arg...)
	if ret, ok = res.(bool); !ok {
		panic(NotExpectedReturn(ret, "boolean").Error())
	}
	return ret
}
