package btree

// BNode represents a node in a B-Tree. It can be an internal or leaf node.
type BNode struct {
    Keys      []string  // Stores keys in the node
    Values    []string  // Stores values in leaf nodes
    Children  []*BNode  // Child nodes (used only in internal nodes)
    IsLeaf    bool      // True if the node is a leaf node
}

// AddKeyValue adds a key-value pair to a leaf node.
func (node *BNode) AddKeyValue(key, value string) {
    node.Keys = append(node.Keys, key)
    node.Values = append(node.Values, value)
}

// SplitNode splits a node if it overflows, creating a new sibling node.
func (node *BNode) SplitNode() *BNode {
    mid := len(node.Keys) / 2

    // New sibling node
    sibling := &BNode{
        Keys:     node.Keys[mid:],
        Values:   node.Values[mid:],
        Children: node.Children[mid:],
        IsLeaf:   node.IsLeaf,
    }

    // Adjust current node keys and children
    node.Keys = node.Keys[:mid]
    node.Values = node.Values[:mid]
    if !node.IsLeaf {
        node.Children = node.Children[:mid]
    }

    return sibling
}
