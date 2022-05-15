package lib

/* Generic single value state with Setter, Getter and ability to add handler, and notify all listeners  */
type GenericState[T any] struct {
	handlers []func(T)
	value    T
}

func NewGenericState[T any](value T) *GenericState[T] {
	state := &GenericState[T]{
		value: value,
	}
	return state
}

func (s *GenericState[T]) SetState(value T) {
	s.value = value
	s.NotifyAll()
}

func (s *GenericState[T]) GetState() T {
	return s.value
}

func (s *GenericState[T]) AddHandler(handler func(T)) {
	s.handlers = append(s.handlers, handler)
}

func (s *GenericState[T]) NotifyAll() {
	for _, h := range s.handlers {
		h(s.value)
	}
}
