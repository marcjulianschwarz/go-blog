package tag

import (
	"fmt"
	"sort"
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

func SortTagsByName(tags []*Tag, descending bool) []*Tag {
	sort.Slice(tags, func(i, j int) bool {
		// TODO: ToLower calls can be reduced here
		if descending {
			return strings.ToLower(tags[i].Name) > strings.ToLower(tags[j].Name)
		}
		return strings.ToLower(tags[i].Name) < strings.ToLower(tags[j].Name)
	})
	return tags
}

func TagToString(tag Tag) string {
	return tag.Name
}

// Converts a list of tags into a comma separated string
func TagsToString(tags []Tag) string {
	tagStrings := make([]string, 0, len(tags))
	for _, tag := range tags {
		tagStrings = append(tagStrings, TagToString(tag))
	}
	return strings.Join(tagStrings, ",")
}
