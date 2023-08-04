package tailo

import (
	"fmt"
	"strings"
)

type writer int

func (w writer) Write(p []byte) (n int, err error) {
	content := strings.TrimSpace(string(p))
	if content == "" {
		return len(p), nil
	}

	fmt.Println("[tailo]", content)
	return len(p), nil
}
