package observer

type NumberGetter interface {
	Number() int
}

type NumberGenerator struct {
	observers []Observer
}

func (ng *NumberGenerator) AddObserver(observer Observer) {
	ng.observers = append(ng.observers, observer)
}

func (ng *NumberGenerator) DeleteObserver(observer Observer) {
	ng.observers = remove(ng.observers, observer)
}

func remove(observers []Observer, observer Observer) []Observer {
	result := []Observer{}
	for _, v := range observers {
		if v != observer {
			result = append(result, v)
		}
	}

	return result
}
