package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

// Build processes a slice of input Records with only and ID and a parent ID into a
// tree structure of Nodes with a slice of children
func Build(input []Record) (*Node, error) {
	if len(input) == 0 {
		return nil, nil
	}
	sort.Slice(input, func(i, j int) bool {
		return input[i].ID < input[j].ID
	})
	if input[0].ID != 0 {
		return nil, errors.New("Root node with ID 0 is required.")
	}
	if input[0].Parent != 0 {
		return nil, errors.New("Root node must have itself as parent.")
	}
	nodesByID := make(map[int]*Node)
	for i, r := range input {
		if r.ID != i {
			return nil, errors.New("Inputs are non-continuous.")
		}
		var newNode *Node
		if n, ok := nodesByID[r.ID]; ok {
			newNode = n
		} else {
			newNode = &Node{ID: r.ID}
			nodesByID[r.ID] = newNode
		}
		if r.ID != 0 {
			if r.Parent > r.ID || r.Parent == r.ID {
				return nil, errors.New("Parent ID must be lower than node ID.")
			}
			if _, ok := nodesByID[r.Parent]; !ok {
				newParent := Node{
					ID: r.Parent,
				}
				nodesByID[r.Parent] = &newParent
			}
			p := nodesByID[r.Parent]
			p.Children = append(p.Children, newNode)
		}
	}
	return nodesByID[0], nil
}
