package tree

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

//Build a tree given an array of records
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	if len(records) == 1 {
		return *Node{
			records[0].ID,
			nil,
		}, nil
	}
	return nil, nil
}
