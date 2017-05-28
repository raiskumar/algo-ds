package common

type queue struct {
	values []*Node
}

func New() *queue {
	queue := &queue{}
	return queue
}

func (q *queue) enqueue(node *Node) {
	q.values = append(q.values, node)
}

func (q *queue) dequeue() *Node {
	var val Node
	if q.isEmpty() {
		return nil
	}
	val = *q.values[0]
	q.values = q.values[1:]
	return &val
}

func (q *queue) isEmpty() bool {
	return len(q.values) == 0
}
