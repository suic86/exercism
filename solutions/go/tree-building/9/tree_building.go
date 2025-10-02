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

// Build function generates a forum tree from a given unsorted set of records.
func Build(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	nodes := make([]*Node, len(records))

	for i, r := range records {

		if i != r.ID || r.Parent > r.ID || r.ID > 0 && r.Parent == r.ID {
			return nil, errors.New("invalid input data")
		}

		nodes[r.ID] = &Node{ID: r.ID}

		if r.ID != 0 {
			nodes[r.Parent].Children = append(nodes[r.Parent].Children, nodes[r.ID])
		}

	}

	return nodes[0], nil
}
