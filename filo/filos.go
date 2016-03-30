package filo

var has filos

func init() {
	has = make(filos)
	Set(builtIns...)
}

type filos map[string]Filo

func Get(key string) Filo {
	if f, exists := has[key]; exists {
		return f
	}
	return NilFilo
}

func Set(fs ...Filo) {
	for _, f := range fs {
		has[f.Key()] = f
	}
}
