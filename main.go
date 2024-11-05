package main

import (
	"fmt"

	"github.com/tacheraSasi/go-minidb.git~/btree"
	"github.com/tacheraSasi/go-minidb.git~/kvstore"
)

func main() {
    // Initializing KV Store
    kv := kvstore.KVStore{Tree: &btree.BTree{}}

    // Set and Get operations
    if err := kv.Set("username", "johndoe"); err != nil {
        fmt.Println("Set Error:", err)
    }

    value, err := kv.Get("username")
    if err != nil {
        fmt.Println("Get Error:", err)
    } else {
        fmt.Println("Retrieved Value:", value)
    }
}
