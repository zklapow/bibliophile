package main

import (
    "github.com/zklapow/bibliophile/models"
    "time"
    "github.com/pquerna/ffjson/ffjson"
    "fmt"
)

func main() {
    book := &models.Book{Created: time.Now().Unix(), Title: "test", Author: "testAuthor", Finished: false}

    data, err := ffjson.Marshal(book)
    if err != nil {
        fmt.Errorf("Error encoding book: %v", err)
        return
    }

    fmt.Println("Encoded book to JSON: %v", string(data))
}
