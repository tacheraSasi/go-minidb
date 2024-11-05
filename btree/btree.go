package btree

// BTree represents a B-Tree data structure for indexing keys.
type BTree struct {
    Root *BNode // Root node of the tree
}

// Insert adds a key-value pair to the B-Tree.
func (tree *BTree) Insert(key, value string) {
    if tree.Root == nil {
        tree.Root = &BNode{IsLeaf: true}
    }

    // Insert key-value pair into the root node
    tree.Root.AddKeyValue(key, value)

    // Split the root if it's full
    if len(tree.Root.Keys) > 2 { // Arbitrary max keys per node for simplicity
        sibling := tree.Root.SplitNode()
        newRoot := &BNode{
            Keys:     []string{sibling.Keys[0]},
            Children: []*BNode{tree.Root, sibling},
            IsLeaf:   false,
        }
        tree.Root = newRoot
    }
}

// Find searches for a key in the B-Tree and returns the associated value.
func (tree *BTree) Find(key string) (string, bool) {
    current := tree.Root
    for current != nil {
        for i, k := range current.Keys {
            if key == k {
                if current.IsLeaf {
                    return current.Values[i], true
                }
                break
            }
        }

        if current.IsLeaf {
            return "", false
        }

        // Descend into the appropriate child
        found := false
        for i, k := range current.Keys {
            if key < k {
                current = current.Children[i]
                found = true
                break
            }
        }
        if !found {
            current = current.Children[len(current.Children)-1]
        }
    }
    return "", false
}
