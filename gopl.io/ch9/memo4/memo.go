package memo

import "sync"

//improve performance with duplication reduction
// sometimes the slow function is called twice from the cache
//both consult the cache and see nothing so both call function having one overwrite the other

type entry struct {
	res   result
	ready chan struct{} //closed when res is ready
}

type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex        // guards cache
	cache map[string]*entry //change to use entry that has result AND ready chan
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		//first request, nothing in cache
		//this goroutine computes
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready) //broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()

		<-e.ready //wait for ready condition
	}
	return e.res.value, e.res.err
}
