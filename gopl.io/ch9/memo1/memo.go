package memo

import (
	"io/ioutil"
	"net/http"
)

//building a concurrent non-blocking cache from 9.7
//memoizing a function, caching the result of a function, so its only computed once
//concurrent-safe and avoid contention with single lock designs

type result struct {
	value interface{}
	err   error
}

//Func is the type of the function to memoize
type Func func(key string) (interface{}, error)

type Memo struct {
	f     Func
	cache map[string]result
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//NOTE: not concurrency-safe
func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}
