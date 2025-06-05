package tag

import (
	"fmt"
	"strings"
)

type Tag struct {
	ID    string
	Name  string
	Color string
	URL   string
}

func (t Tag) String() string {
	return fmt.Sprintf("Tag %s", t.Name)
}

func TagNameToId(name string) (id string) {
	return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}
