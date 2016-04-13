package extension

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
