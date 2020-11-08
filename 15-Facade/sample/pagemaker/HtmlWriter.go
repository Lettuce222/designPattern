package pagemaker

import (
	"fmt"
	"os"
)

type HtmlWriter struct {
	file *os.File
}

func NewHtmlWriter(filename string) *HtmlWriter {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("エラー処理をサボりました %v", err)
	}

	return &HtmlWriter{file: file}
}

func (hw *HtmlWriter) title(title string) {
	hw.file.WriteString("<html>\n")
	hw.file.WriteString("<head>\n")
	hw.file.WriteString(fmt.Sprintf("<title>%v</title>\n", title))
	hw.file.WriteString("</head>\n")
	hw.file.WriteString("<body>\n")
	hw.file.WriteString(fmt.Sprintf("<h1>%v</h1>\n", title))
}

func (hw *HtmlWriter) paragraph(msg string) {
	hw.file.WriteString(fmt.Sprintf("<p>%v</p>\n", msg))
}

func (hw *HtmlWriter) link(href string, caption string) {
	hw.file.WriteString(fmt.Sprintf("<a href=%v>%v</p>\n", href, caption))
}

func (hw *HtmlWriter) mailto(mailaddr string, username string) {
	hw.link(fmt.Sprintf("mailto:%v", mailaddr), username)
}

func (hw *HtmlWriter) close() {
	hw.file.WriteString("</body>\n")
	hw.file.WriteString("</html>\n")
	hw.file.Close()
}
