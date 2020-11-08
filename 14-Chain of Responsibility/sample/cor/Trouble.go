package cor

import (
	"fmt"
)

type Trouble struct {
	Number int
}

func (t *Trouble) string() string {
	return fmt.Sprintf("[Trouble %v]", t.Number)
}
