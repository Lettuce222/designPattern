package pagemaker

import (
	"fmt"
)

func MakeWelcomePage(mailaddr string, filename string) {
	mailprop := getProperties("maildata")
	username := mailprop[mailaddr]
	writer := NewHtmlWriter(filename)

	writer.title(fmt.Sprintf("Welcome to %v's page", username))
	writer.paragraph(fmt.Sprintf("%vのページへようこそ", username))
	writer.paragraph("メール待っていますね")
	writer.mailto(mailaddr, username)
	writer.close()
	fmt.Printf("%v is created for %v(%v)\n", filename, mailaddr, username)
}
