//go:generate ffjson book.go

package models

type Book struct {
    Id         int64;
    Created    int64;

    Title      string;
    Author     string;
    Genre      string;

    Finished   bool;
    FinishedAt int64;
}
