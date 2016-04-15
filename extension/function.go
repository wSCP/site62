package extension

type Functions interface {
	Add(...Function)
	Functions() []Function
}

type functions struct {
	functions []Function
}

func (f *functions) Add(fns ...Function) {
	f.functions = append(f.functions, fns...)
}

func (f *functions) Functions() []Function {
	return f.functions
}

type Function interface {
	Key() string
	Value() interface{}
}

type function struct {
	key   string
	value interface{}
}

func NewFunction(key string, fn interface{}) Function {
	return &function{
		key:   key,
		value: fn,
	}
}

func (f *function) Key() string {
	return f.key
}

func (f *function) Value() interface{} {
	return f.value
}
