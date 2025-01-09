package kata

func FindLoop(node Node) int {
	loopNode := detectLoop(node)
	return loopSize(loopNode)
}

func detectLoop(node Node) Node {
	slow := node
	fast := node.Next()

	for slow.Value() != fast.Value() {
		slow = slow.Next()
		fast = fast.Next().Next()
	}

	return slow
}

func loopSize(node Node) int {
	loopNode := node
	length := 1

	for loopNode.Next().Value() != node.Value() {
		loopNode = loopNode.Next()
		length++
	}

	return length
}
