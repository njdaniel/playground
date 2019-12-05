package main

func main() {
	
}


type Logger struct {
	m map[string]int
}


/** Initialize your data structure here. */
func Constructor() Logger {
	logger := &Logger{
		m: make(map[string]int, 0),
	}
	return *logger
}


/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
  If this method returns false, the message will not be printed.
  The timestamp is in seconds granularity. */
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	var updateHash = func() {
		this.m[message] = timestamp
	}
	if _, ok := this.m[message]; ok {
		if timestamp-10 >= this.m[message] {
			updateHash()
			return true
		} else {
			return false
		}
	}
	updateHash()
	return true
}


/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */