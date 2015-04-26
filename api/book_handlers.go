package api
import (
    "github.com/zklapow/bibliophile/persist"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/Sirupsen/logrus"
    "github.com/pquerna/ffjson/ffjson"
    "strconv"
)

type BookHandler struct {
    Books *persist.Books `inject:""`
    Logger *logrus.Logger
}

func (h *BookHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    id, err := strconv.ParseInt(vars["id"], 0, 64)
    if err != nil {
        resp.WriteHeader(http.StatusBadRequest)
        resp.Write([]byte("ID must be an int"))
        return
    }

    log := h.Logger.WithField("book", id)

    book, err := h.Books.Get(id)
    if err != nil {
        resp.WriteHeader(http.StatusInternalServerError)
        log.Errorf("Error fetching book %v", err)
        return
    }

    json, err := ffjson.Marshal(book)
    if err != nil {
        resp.WriteHeader(http.StatusInternalServerError)
        log.Errorf("Error marshalling book: %v", err)
        return
    }

    resp.Header().Add("Content-Type", "application/json")
    resp.Write(json)
}