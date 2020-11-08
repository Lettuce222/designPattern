package observer

type Observer interface {
	update(generator NumberGenerator)
}
