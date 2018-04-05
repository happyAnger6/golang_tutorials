package chapter9

type request struct {
	key string
	response chan <- result
}

type Memo5 struct {
	requests chan request
	f Func
}

type entry struct {
	res result
	ready chan struct{}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan <- result) {
	<- e.ready
	response <- e.res
}

func NewMemo5(f Func) *Memo5 {
	memo := &Memo5{ requests: make(chan request),
	f: f}

	go memo.server()
	return memo
}

func (m *Memo5) Get(key string) (interface{}, error) {
	response := make(chan result)
	m.requests <- request{key, response}
	res := <- response
	return res.value, res.err
}

func (m *Memo5)Close() {
	close(m.requests)
}

func (m *Memo5)server() {
	cache := make(map[string]*entry)
	for req := range m.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(m.f, req.key)
		}
		go e.deliver(req.response)
	}
}




