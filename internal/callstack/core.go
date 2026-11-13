package callstack

import (
	"sync"

	"github.com/qean32/holo/internal/model"
)

var (
	mu      sync.Mutex
	queue   []model.Event
	changed = make(chan struct{}, 1)
)

func Push(event model.Event) {
	mu.Lock()
	queue = append(queue, event)
	mu.Unlock()

	select {
	case changed <- struct{}{}:
	default:
	}
}

func PushMany(events ...model.Event) {
	if len(events) == 0 {
		return
	}

	mu.Lock()
	queue = append(queue, events...)
	mu.Unlock()

	select {
	case changed <- struct{}{}:
	default:
	}
}

func Pop() (model.Event, bool) {
	mu.Lock()
	defer mu.Unlock()

	if len(queue) == 0 {
		return model.Event{}, false
	}

	last := len(queue) - 1
	event := queue[last]
	queue[last] = model.Event{}
	queue = queue[:last]
	return event, true
}

func Shift() (model.Event, bool) {
	mu.Lock()
	defer mu.Unlock()

	if len(queue) == 0 {
		return model.Event{}, false
	}

	event := queue[0]
	queue[0] = model.Event{}
	queue = queue[1:]
	return event, true
}

func Changed() <-chan struct{} {
	return changed
}
