package node

type Treehead interface {
	Head() Node
	SetHead(Node)
}

type Treetail interface {
	Tail() []Node
	SetTail(...Node)
}

type Tree interface {
	Treehead
	Treetail
}

type tree struct {
	head Node
	tail []Node
}

var EmptyTree tree = tree{nil, make([]Node, 0)}

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
