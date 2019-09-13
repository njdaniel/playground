package memo

//memo3 aims to improve the performance that was lost
//from memo2 trying to make it concurrency-safe

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.f(key)
	if !ok {
		res.value, res.err = memo.f(key)
		memo.mu.Lock()
		memo.cache[key] =resmemo.mu.Unlock()
	}
	return res.value, res.err
}
