package tree

import (
	"fmt"
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

func (r Record) String() string {
	return fmt.Sprintf("Record{ID: %d, Parent: %d}", r.ID, r.Parent)
}

// Build function generates a forum tree from a given unsorted set of records.
func Build(records []Record) (*Node, error) {
	recordCount := len(records)

	if recordCount == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	nodes := make([]*Node, recordCount, recordCount)

	for i, r := range records {

		id, parent := r.ID, r.Parent

		if id == 0 && parent != 0 {
			return nil, fmt.Errorf("root node mustn't have a parent node other then itself: %s", r)
		}
		if id != 0 && id == parent {
			return nil, fmt.Errorf("non-root node parent to itself: %s", r)
		}
		if id >= recordCount {
			return nil, fmt.Errorf("non-contiguous nodes: %s (ID >= number of records: %d >= %d)", r, r.ID, len(records))
		}
		if parent > id {
			return nil, fmt.Errorf("child node has lower ID then the parent node: %s", r)
		}
		if i != id {
			return nil, fmt.Errorf("duplicate id: %d", r.ID)
		}

		nodes[id] = &Node{ID: id}

		if id != 0 {
			nodes[parent].Children = append(nodes[parent].Children, nodes[id])
		}
	}

	return nodes[0], nil
}
