package global

import (
	"sync/atomic"
)

var scenario atomic.Bool

func ScenarioEnabled() bool {
	return scenario.Load()
}

func ScenarioToggle() bool {
	for {
		cur := scenario.Load()
		if scenario.CompareAndSwap(cur, !cur) {
			return !cur
		}
	}
}
