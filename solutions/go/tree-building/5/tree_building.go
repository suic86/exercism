package tree

import (
	"errors"
	"sort"
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

// ByID implements sort.Interface for []Record based on the ID field.
type ByID []Record

func (b ByID) Len() int           { return len(b) }
func (b ByID) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByID) Less(i, j int) bool { return b[i].ID < b[j].ID }

// Build function generates a forum tree from a given unsorted set of records.
func Build(records []Record) (*Node, error) {
	recordCount := len(records)

	if recordCount == 0 {
		return nil, nil
	}

	sort.Sort(ByID(records))

	// Adjancency list
	adjacent := make([]int, recordCount, recordCount)
	for id := 0; id < recordCount; id++ {
		adjacent[id] = -1
	}

	nodes := make([]*Node, recordCount, recordCount)

	// Result tree
	nodes[0] = &Node{ID: 0}

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

		if id != 0 {
			nodes[id] = &Node{ID: id}
			nodes[parent].Children = append(nodes[parent].Children, nodes[id])
		}
	}

	return nodes[0], nil
}
