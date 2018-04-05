package chapter9

type result struct {
	value interface{}
	err   error
}

type Func func(key string) (interface{}, error)

type Memo struct {
	f     Func
	cache map[string]result
}

func New(f Func) *Memo {
	return &Memo{f: f,
		cache: make(map[string]result)}
}

func (m *Memo) Get(key string) (interface{}, error) {
	v, ok := m.cache[key]
	if !ok {
		v.value, v.err = m.f(key)
		m.cache[key] = v
	}
	return v.value, v.err
}
