package api

import (
    "github.com/zklapow/bibliophile/persist"
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
    "github.com/zklapow/bibliophile/models"
)

type BookHandler struct {
    JsonRestHttpHandler
    Books *persist.Books `inject:""`
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

func NewBookHandler() *BookHandler {
    handler := &BookHandler{}
    handler.JsonRestHttpHandler.Handler = handler
    return handler
}
