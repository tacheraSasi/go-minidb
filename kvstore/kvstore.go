package kvstore

import (
	"os"

	"github.com/tacheraSasi/go-minidb.git~/btree"
)

// KVStore combines a BTree and persistence for a KV database.
type KVStore struct {
    Tree *btree.BTree
}

func (kv *KVStore) Set(key, value string) error {
    // Add key-value pair to B-Tree for indexing
    kv.Tree.Insert(key, value)

    // Persist to disk
    path := "data/" + key
    return SaveData(path, []byte(value))
}


func (kv *KVStore) Get(key string) (string, error) {
    path := "data/" + key
    data, err := os.ReadFile(path)
    if err != nil {
        return "", err
    }
    return string(data), nil
}
