package main

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	ID       int
	PID      int
	Children []*Node
}

func main() {
	var nodes = []*Node{
		{ID: 1, PID: 0},
		{ID: 2, PID: 1},
		{ID: 3, PID: 1},
		{ID: 4, PID: 2},
		{ID: 5, PID: 2},
		{ID: 6, PID: 3},
	}

	tree := buildTree(nodes)

	b, err := json.Marshal(tree)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func buildTree(nodes []*Node) *Node {
	if len(nodes) == 0 {
		return nil
	}

	nodeMap := make(map[int]*Node)
	for _, node := range nodes {
		nodeMap[node.ID] = node
	}

	var root *Node
	for _, node := range nodes {
		if node.PID == 0 {
			root = node
		} else {
			parent, ok := nodeMap[node.PID]
			if !ok {
				continue
			}
			parent.Children = append(parent.Children, node)
		}
	}

	return root
}
