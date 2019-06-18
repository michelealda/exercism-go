package tree

import (
	"fmt"
	"sort"
)

//Record is a unstructered tree entry
type Record struct {
	ID     int
	Parent int
}

//Node represent a tree node
type Node struct {
	ID       int
	Children []*Node
}

//NodeSlice is a list of nodes
type NodeSlice []*Node

func (n NodeSlice) Len() int           { return len(n) }
func (n NodeSlice) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n NodeSlice) Less(i, j int) bool { return n[i].ID < n[j].ID }

//Build a tree given an array of records
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	nodes := make([]Node, len(records))
	parents := make([]*Node, len(records))
	visited := make([]bool, len(records))

	for _, record := range records {
		if record.ID >= len(records) {
			return nil, fmt.Errorf("Id %d out of range", record.ID)
		}
		if record.ID != 0 && record.ID <= record.Parent {
			return nil, fmt.Errorf("Invalid record %d, %d", record.ID, record.Parent)
		}
		if record.ID == 0 && record.Parent != 0 {
			return nil, fmt.Errorf("Invalid record %d, %d", record.ID, record.Parent)
		}
		if visited[record.ID] {
			return nil, fmt.Errorf("Duplicate %d", record.ID)
		}

		visited[record.ID] = true
		if record.ID != 0 {
			parents[record.ID] = &nodes[record.Parent]
		}
	}

	for i := 1; i < len(nodes); i++ {
		parents[i].Children = append(parents[i].Children, &nodes[i])
	}

	for i, node := range nodes {
		nodes[i].ID = i
		sort.Sort(NodeSlice(node.Children))
	}
	return &nodes[0], nil
}
