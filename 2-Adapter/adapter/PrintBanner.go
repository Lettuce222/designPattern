package adapter

type PrintBanner struct {
	banner *Banner
}

func NewPrintBanner(String string) *PrintBanner {
	return &PrintBanner{&Banner{String: String}}
}

func (pb *PrintBanner) PrintWeak() {
	pb.banner.showWithParen()
}

func (pb *PrintBanner) PrintStrong() {
	pb.banner.showWithAster()
}
