package emitter

// Emitter struct
type Emitter struct {
	listeners map[string][]func(...interface{})
}

// On func
func (e *Emitter) On(event string, listener func(...interface{})) {
	if events, ok := e.listeners[event]; ok {
		e.listeners[event] = append(events, listener)
	}
}

// Emit func
func (e Emitter) Emit(event string, data ...interface{}) {
	if events, ok := e.listeners[event]; ok {
		for _, listener := range events {
			listener(data...)
		}
	}
}
