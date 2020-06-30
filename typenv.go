package typenv

import "sync"

var mutex sync.Mutex
var registry map[string]bool

func init() {
	registry = map[string]bool{}
}

func use(name string) {
	if _, ok := registry[name]; !ok {
		mutex.Lock()
		registry[name] = true
		mutex.Unlock()
	}
}
