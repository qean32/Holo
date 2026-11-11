package global

import (
	"sync/atomic"
)

var script atomic.Bool

func ScriptEnabled() bool {
	return script.Load()
}

func ScriptToggle() bool {
	for {
		cur := script.Load()
		if script.CompareAndSwap(cur, !cur) {
			return !cur
		}
	}
}
