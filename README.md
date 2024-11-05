# KV Store

This is a simple Key-Value store built from scratch in Go, using B-Trees for indexing and atomic file persistence.

## Project Structure
- **`data/`**: Stores key-value data as files.
- **`btree/`**: Implements B-Tree for efficient indexing.
- **`kvstore/`**: Contains main KV store logic and persistence functions.
- **`main.go`**: Entry point with example usage.

## Usage
To run the KV store:

1. Initialize the project:
    ```bash
    go mod init kv-store
    ```

2. Run the main program:
    ```bash
    go run main.go
    ```

3. Expected output:
    ```
    Retrieved Value: johndoe
    ```

This code demonstrates setting and getting a key-value pair, ensuring data is stored persistently.
