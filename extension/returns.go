package extension

type Returns interface {
	RunString(string, ...interface{}) string
	RunInteger(string, ...interface{}) int
	RunBoolean(string, ...interface{}) bool
}

type returns struct {
	e Extension
}

func (r *returns) RunString(name string, arg ...interface{}) string {
	var ret string
	var ok bool
	res := r.e.MustRun(name, arg...)
	if ret, ok = res.(string); !ok {
		panic(NotExpectedReturn(ret, "string").Error())
	}
	return ret
}

func (r *returns) RunInteger(name string, arg ...interface{}) int {
	var ret int
	var ok bool
	res := r.e.MustRun(name, arg...)
	if ret, ok = res.(int); !ok {
		panic(NotExpectedReturn(ret, "integer").Error())
	}
	return ret
}

func (r *returns) RunBoolean(name string, arg ...interface{}) bool {
	var ret bool
	var ok bool
	res := r.e.MustRun(name, arg...)
	if ret, ok = res.(bool); !ok {
		panic(NotExpectedReturn(ret, "boolean").Error())
	}
	return ret
}
