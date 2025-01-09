package kata

type Node interface {
	Value() int
	Next() Node
}

type simpleNode struct {
	value int
	next  Node
}

func (n simpleNode) Value() int {
	return n.value
}

func (n simpleNode) Next() Node {
	return n.next
}

func newNode(value int, next Node) *simpleNode {
	return &simpleNode{
		value: value,
		next:  next,
	}
}

type chain struct {
	head  *simpleNode
	nodes map[int]*simpleNode
}

func (c *chain) addNode(value, next int) Node {
	node, ok := c.nodes[value]
	if !ok {
		node = newNode(value, nil)
		c.nodes[value] = node
	}

	nextNode, ok := c.nodes[next]
	if !ok {
		nextNode = newNode(next, nil)
		c.nodes[next] = nextNode
	}

	node.next = nextNode

	if c.head == nil {
		c.head = node
	}

	return node
}

func CreateChain(values map[int]int) Node {
	c := &chain{
		nodes: make(map[int]*simpleNode),
	}

	for value, next := range values {
		c.addNode(value, next)
	}

	return c.head
}

func CreateRandomChain(prefixCount, loopCount uint) Node {
	values := make(map[int]int)

	for i := 0; i < int(prefixCount); i++ {
		values[i] = i + 1
	}

	for i := 0; i < int(loopCount)-1; i++ {
		values[int(prefixCount)+i] = int(prefixCount) + i + 1
	}

	values[int(prefixCount)+int(loopCount)-1] = int(prefixCount)

	return CreateChain(values)
}
