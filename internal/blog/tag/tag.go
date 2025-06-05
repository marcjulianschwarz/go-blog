package tag

import "fmt"

type Tag struct {
	ID    string
	Name  string
	Color string
	URL   string
}

func (t Tag) String() string {
	return fmt.Sprintf("Tag %s", t.Name)
}
