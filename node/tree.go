package node

type Tree interface {
	Head() Node
	SetHead(Node)
	Tail() []Node
	SetTail(...Node)
}

type tree struct {
	head Node
	tail []Node
}

func NewTree(h Node, t ...Node) Tree {
	tr := &tree{
		head: h,
		tail: make([]Node, 0),
	}
	tr.SetTail(t...)
	return tr
}

func (t *tree) Head() Node {
	return t.head
}

func (t *tree) SetHead(n Node) {
	t.head = n
}

func (t *tree) Tail() []Node {
	return t.tail
}

func (t *tree) SetTail(ns ...Node) {
	t.tail = append(t.tail, ns...)
}
