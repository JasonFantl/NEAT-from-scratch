package neat

type nodeType int

const (
	SENSOR nodeType = iota + 1
	HIDDEN
	OUTPUT
)

type Node struct {
	ID    int
	State nodeType
}

func NewNode(id int, state nodeType) Node {
	return Node{id, state}
}
