package main

import (
    "github.com/zklapow/bibliophile/models"
    "time"
    "github.com/pquerna/ffjson/ffjson"
    "github.com/garyburd/redigo/redis"
    "flag"
    "github.com/facebookgo/inject"
    "github.com/zklapow/bibliophile/persist"
    "os"
    "github.com/Sirupsen/logrus"
)

var (
    pool *redis.Pool
    log = logrus.New()
    redisServer = flag.String("redis", ":6379", "Address of redis server.")
)

func main() {
    var g inject.Graph

    pool = buildPool(*redisServer)
    var books persist.Books

    err := g.Provide(&inject.Object{Value: pool}, &inject.Object{Value: &books})
    if err != nil {
        log.Errorf("Error building dep graph: %v", err)
        os.Exit(1)
    }

    if err = g.Populate(); err != nil {
        log.Errorf("Error populating graph: %v", err)
        os.Exit(1)
    }

    book := &models.Book{Created: time.Now().Unix(), Title: "test", Author: "testAuthor", Finished: false}

    data, err := ffjson.Marshal(book)
    if err != nil {
        log.Errorf("Error encoding book: %v", err)
        return
    }

    log.Infof("Encoded book to JSON: %v", string(data))
}

func buildPool(server string) *redis.Pool {
    return &redis.Pool{
        MaxIdle: 3,
        IdleTimeout: 240 * time.Second,
        Dial: func () (redis.Conn, error) {
            c, err := redis.Dial("tcp", server)
            if err != nil {
                return nil, err
            }
            return c, err
        },
        TestOnBorrow: func(c redis.Conn, t time.Time) error {
            _, err := c.Do("PING")
            return err
        },
    }
}

func upsertBook(book *models.Book) error {
    return nil
}
