package utils

import "fmt"

type Node struct {
	Neighbors []*Node
}

type Graph struct {
	nodes map[int]*Node
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[int]*Node),
	}
}

func (g *Graph) AddNode(nodeID int) {
	if _, exists := g.nodes[nodeID]; !exists {
		newNode := &Node{
			Neighbors: []*Node{},
		}

		g.nodes[nodeID] = newNode

		fmt.Println("new node added to graph")
	} else {
		fmt.Println("node already exists")
	}
}

func (g *Graph) AddEdge(nodeID1, nodeID2 int) {
	node1 := g.nodes[nodeID1]
	node2 := g.nodes[nodeID2]

	node1.Neighbors = append(node1.Neighbors, node2)
	node2.Neighbors = append(node2.Neighbors, node1)
}

func (g *Graph) removeEdge(node, neighbor *Node) {
	index := -1

	for i, n := range node.Neighbors {
		if n == neighbor {
			index = i
			break
		}
	}

	if index != -1 {
		node.Neighbors =
			append(node.Neighbors[:index], node.Neighbors[index+1:]...)
	}
}

func (g *Graph) RemoveEdge(node, neighbor *Node) {
	g.removeEdge(node, neighbor)
	g.removeEdge(neighbor, node)
}

func (g *Graph) RemoveNode(nodeID int) {
	node, exists := g.nodes[nodeID]
	if !exists {
		fmt.Println("node does not exist")
		return
	}

	for _, neighbor := range node.Neighbors {
		g.RemoveEdge(node, neighbor)
	}

	delete(g.nodes, nodeID)
	fmt.Println("node deleted")
}
