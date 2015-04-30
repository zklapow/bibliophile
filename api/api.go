package api

import (
    "github.com/facebookgo/inject"
    "github.com/zklapow/bibliophile/persist"
    "github.com/Sirupsen/logrus"
    "os"
    "github.com/gorilla/mux"
    "github.com/zklapow/negronilogrus"
    "github.com/pilu/xrequestid"
    "github.com/codegangsta/negroni"
    "fmt"
    "github.com/zklapow/bibliophile/util"
    "gopkg.in/tylerb/graceful.v1"
    "time"
)

var (
    logger = logrus.New()
)

func Start(port int, redisServer string) {
    var g inject.Graph

    pool := util.BuildPool(redisServer)
    var books persist.Books
    bookHandler := &BookHandler{}
    booksHandler := &BooksHandler{}

    err := g.Provide(&inject.Object{Value: pool},
                     &inject.Object{Value: &books},
                     &inject.Object{Value: bookHandler},
                     &inject.Object{Value: booksHandler},
                     &inject.Object{Value: logger})

    if err != nil {
        logger.Errorf("Error building dep graph: %v", err)
        os.Exit(1)
    }

    if err = g.Populate(); err != nil {
        logger.Errorf("Error populating graph: %v", err)
        os.Exit(1)
    }

    m := mux.NewRouter()
    router := m.PathPrefix("/bibliophile/v1").Subrouter()
    router.HandleFunc("/books/{id}", requestHandler(bookHandler))
    router.HandleFunc("/books", requestHandler(booksHandler))

    n := negroni.New(negroni.NewRecovery(), xrequestid.New(16), negronilogrus.NewMiddleware())
    n.UseHandler(m)

    logrus.Infof("Starting server on port %v", port)
    graceful.Run(fmt.Sprintf(":%v", port), 10*time.Second, n)
}
