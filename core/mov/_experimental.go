package mov

import (
	"strings"
)

// Path is designed to retrieve  ID  from the URI path
// We were not recommended to use frameworks like gorilla, right?
type Path struct {
	ID   string
	Path string
}

const Slash = "/"

func NewPath(path string) *Path {
	var id string
	path = strings.Trim(path, Slash)
	str := strings.Split(path, Slash)
	if len(str) > 1 {
		id = str[len(str)-1]
		path = strings.Join(str[:len(str)-1], Slash)
	}

	return &Path{ID: id, Path: path}
}

func (path *Path) HasID() bool {
	return len(path.ID) > 0
}
