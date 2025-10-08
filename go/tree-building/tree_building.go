package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	n := len(records)
	if n == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	for i, r := range records {
		if r.ID != i {
			return nil, errors.New("non-continuous ids or duplicate")
		}
	}

	nodes := make([]*Node, n)
	for _, r := range records {
		if r.ID == r.Parent && r.ID != 0 {
			return nil, errors.New("non-root node points to itself")
		}
		if r.ID != 0 && r.Parent >= r.ID {
			return nil, errors.New("parent id must be lower than child id")
		}
		nodes[r.ID] = &Node{ID: r.ID}
		if r.ID != 0 {
			parent := nodes[r.Parent]
			if parent == nil {
				return nil, errors.New("missing parent")
			}
			parent.Children = append(parent.Children, nodes[r.ID])
		}
	}

	for _, n := range nodes {
		sort.Slice(n.Children, func(i, j int) bool {
			return n.Children[i].ID < n.Children[j].ID
		})
	}

	return nodes[0], nil
}
