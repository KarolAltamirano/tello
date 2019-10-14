package emitter

// Emitter struct
type Emitter struct {
	listeners map[string][]func(...interface{})
}

// NewEmitter Create New Emitter
func NewEmitter() Emitter {
	return Emitter{
		listeners: make(map[string][]func(...interface{}), 5),
	}
}

// On func
func (e *Emitter) On(event string, listener func(...interface{})) {
	e.listeners[event] = append(e.listeners[event], listener)
}

// Emit func
func (e Emitter) Emit(event string, data ...interface{}) {
	if events, ok := e.listeners[event]; ok {
		for _, listener := range events {
			listener(data...)
		}
	}
}
