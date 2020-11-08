package observer

type NumberExecuter interface {
	Number() int
	innerExecute(NotifyObservers func())
}

type NumberGenerator struct {
	Observers []Observer
	NumberExecuter
}

func (ng *NumberGenerator) AddObserver(observer Observer) {
	ng.Observers = append(ng.Observers, observer)
}

func (ng *NumberGenerator) DeleteObserver(observer Observer) {
	ng.Observers = remove(ng.Observers, observer)
}

func (ng *NumberGenerator) NotifyObservers() {
	for _, v := range ng.Observers {
		v.update(*ng)
	}
}

func (ng *NumberGenerator) Execute() {
	ng.innerExecute(ng.NotifyObservers)
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
