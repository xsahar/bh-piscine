package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	newNodeL := &NodeL{Data: data, Next: nil}

	if l.Head == nil {
		l.Head = newNodeL
	} else {
		l.Tail.Next = newNodeL
	}
	l.Tail = newNodeL
}
