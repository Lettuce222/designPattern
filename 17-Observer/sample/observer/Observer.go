package observer

type Observer interface {
	update(ng NumberGetter)
}
