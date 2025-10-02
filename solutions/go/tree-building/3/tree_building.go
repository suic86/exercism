package tree

import (
	"errors"
)

// Node represents a single node in the forum tree.
type Node struct {
	ID       int
	Children []*Node
}

// Record stores a single a forum record with reference to its parent record.
type Record struct {
	ID     int
	Parent int
}

// Build function generates a forum tree from a given unsorted set of records.
func Build(records []Record) (*Node, error) {
	recordCount := len(records)

	if recordCount == 0 {
		return nil, nil
	}

	// Construct adjacency list

	adjacent := make([]int, recordCount, recordCount)

	for id := 0; id < recordCount; id++ {
		adjacent[id] = -1
	}

	for _, r := range records {

		id, parent := r.ID, r.Parent

		if id == 0 && parent != 0 {
			return nil, errors.New("root node mustn't have a parent node other then itself")
		}
		if id != 0 && id == parent {
			return nil, errors.New("non-root node parent to itself")
		}
		if id >= recordCount {
			return nil, errors.New("non-contiguous nodes")
		}
		if parent > id {
			return nil, errors.New("child node has lower ID then the parent node")
		}
		if adjacent[id] != -1 {
			if adjacent[id] == parent {
				return nil, errors.New("duplicate node")
			}
			return nil, errors.New("node has multiple parents")
		}

		adjacent[id] = parent
	}

	// Construct the result tree from the adjancency list

	nodes := make([]*Node, recordCount, recordCount)
	nodes[0] = &Node{ID: 0}

	for i, p := range adjacent[1:] {
		nodes[i+1] = &Node{ID: i + 1}
		nodes[p].Children = append(nodes[p].Children, nodes[i+1])
	}

	return nodes[0], nil
}
