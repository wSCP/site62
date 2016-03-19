package filo

var has filos

func init() {
	has = make(filos)
	for _, f := range builtIns {
		Set(f)
	}
}

type filos map[string]Filo

func Get(key string) Filo {
	if f, exists := has[key]; exists {
		return f
	}
	return nil
}

func Set(f Filo) {
	has[f.Key()] = f
}
