package blog

import "fmt"

type Tag struct {
	Name string
}

func (t Tag) String() string {
	return fmt.Sprintf("Tag %s", t.Name)
}
