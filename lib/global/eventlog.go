package global

import "sync/atomic"

var eventLog atomic.Bool

func EventLogEnabled() bool {
	return eventLog.Load()
}

func EventLogToggle() bool {
	for {
		cur := eventLog.Load()
		if eventLog.CompareAndSwap(cur, !cur) {
			return !cur
		}
	}
}
