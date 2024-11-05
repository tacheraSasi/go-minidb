package kvstore

import (
    "fmt"
    "os"
)

// SaveData saves data atomically to a file at the specified path.
func SaveData(path string, data []byte) error {
    // Creating a temporary file path
    tmpPath := fmt.Sprintf("%s.tmp", path)

    // Opening or create the temp file with write permissions
    file, err := os.OpenFile(tmpPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
    if err != nil {
        return err
    }
    defer file.Close()

    // Writing the data to the temp file
    if _, err := file.Write(data); err != nil {
        return err
    }

    // Ensuring data is flushed to disk
    if err := file.Sync(); err != nil {
        return err
    }

    // Renaming the temp file to the target file for atomicity
    return os.Rename(tmpPath, path)
}
