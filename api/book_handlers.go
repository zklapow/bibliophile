package api

import (
    "github.com/zklapow/bibliophile/persist"
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
    "github.com/zklapow/bibliophile/models"
    "io/ioutil"
    "github.com/pquerna/ffjson/ffjson"
    "fmt"
)

type BookHandler struct {
    InjectLogger
    Books *persist.Books `inject:""`

    PostNotAllowed
}

func (h *BookHandler) Get(resp http.ResponseWriter, req *http.Request) (int, interface{}) {
    vars := mux.Vars(req)
    id, err := strconv.ParseInt(vars["id"], 0, 64)
    if err != nil {
        return http.StatusBadRequest, models.NewJsonError("ID must be an int")
    }

    log := h.Log.WithField("book", id)

    book, err := h.Books.Get(id)
    if err != nil {
        log.Errorf("Error fetching book %v", err)
        return http.StatusInternalServerError, nil
    }

    if book == nil {
        return http.StatusNotFound, nil
    }

    return http.StatusOK, book
}

type BooksHandler struct {
    InjectLogger
    Books *persist.Books `inject:""`
}

func (h *BooksHandler) Get(resp http.ResponseWriter, req *http.Request) (int, interface{}) {
    books, err := h.Books.GetAll()
    if err != nil {
        return http.StatusInternalServerError, models.NewJsonError(fmt.Sprintf("Could not get books: %v", err))
    }

    return http.StatusOK, books
}

func (h *BooksHandler) Post(resp http.ResponseWriter, req *http.Request) (int, interface{}) {
    data, err := ioutil.ReadAll(req.Body)
    if err != nil {
        h.Log.Errorf("Error reading JSON: %v", err)
        return http.StatusBadRequest, nil
    }

    var book models.Book
    if err := ffjson.Unmarshal(data, &book); err != nil {
        return http.StatusInternalServerError, models.NewJsonError(fmt.Sprintf("Could nor unmarshal JSON: %v", err))
    }

    if err = h.Books.Create(&book); err != nil {
        return http.StatusInternalServerError, models.NewJsonError(fmt.Sprintf("Error writing to database: %v", err))
    }

    return http.StatusNoContent, book
}
