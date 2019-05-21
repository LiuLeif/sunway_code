// 2018-02-06 16:58
package main

import (
	"os"

	"github.com/bradleyjkemp/memmap"
)

type Node struct {
	label     int
	neighbors []*Node
}

func NewNode(label int) *Node {
	node := Node{}
	node.label = label
	node.neighbors = []*Node{}
	return &node
}

func helper(node *Node, m map[*Node]*Node) *Node {
	if m[node] != nil {
		return m[node]
	}

	copy := NewNode(node.label)
	m[node] = copy
	for _, neighbor := range node.neighbors {
		copy.neighbors = append(copy.neighbors, helper(neighbor, m))
	}
	return copy
}

func cloneGraph(node *Node) *Node {
	m := map[*Node]*Node{}
	return helper(node, m)
}

func main() {
	n1 := NewNode(1)
	n2 := NewNode(2)
	n3 := NewNode(3)

	n1.neighbors = append(n1.neighbors, n2)
	n1.neighbors = append(n1.neighbors, n3)

	n2.neighbors = append(n2.neighbors, n1)
	n2.neighbors = append(n2.neighbors, n3)

	n3.neighbors = append(n3.neighbors, n1)
	n3.neighbors = append(n3.neighbors, n2)

	n := cloneGraph(n1)
	f, _ := os.Create("/tmp/memmap.graph")
	memmap.Map(f, &n)
}
