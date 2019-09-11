package memo

//memo2 fixing the data race for the Get method of memo
//adding a mutex to lock and unlock and cache inside the critical section

import "sync"

// Func is the type of the function to memoize.
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

//!+

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]result
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	memo.mu.Unlock()
	return res.value, res.err
}
