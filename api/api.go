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
)

var (
    log = logrus.New()
)

func Start(port int, redisServer string) {
    var g inject.Graph

    pool := util.BuildPool(redisServer)
    var books persist.Books
    var bookHandler BookHandler

    err := g.Provide(&inject.Object{Value: pool}, &inject.Object{Value: &books}, &inject.Object{Value: &bookHandler})
    if err != nil {
        log.Errorf("Error building dep graph: %v", err)
        os.Exit(1)
    }

    if err = g.Populate(); err != nil {
        log.Errorf("Error populating graph: %v", err)
        os.Exit(1)
    }

    m := mux.NewRouter()
    router := m.PathPrefix("/bibliophile/v1").Subrouter()
    router.Handle("/books/{id}", &bookHandler)

    n := negroni.New(negroni.NewRecovery(), xrequestid.New(16), negronilogrus.NewMiddleware())
    n.UseHandler(m)
    n.Run(fmt.Sprintf(":%v", port))
}
