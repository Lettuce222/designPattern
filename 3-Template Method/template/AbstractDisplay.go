package template

type OpenPrintCloser interface {
	open()
	print()
	close()
}

type AbstractDisplay struct {
	OpenPrintCloser
}

func NewAbstractDisplay(opc OpenPrintCloser) *AbstractDisplay {
	return &AbstractDisplay{opc}
}

func (ab AbstractDisplay) Display() {
	ab.open()
	for i := 0; i < 5; i++ {
		ab.print()
	}
	ab.close()
}
