package persist
import (
    "github.com/garyburd/redigo/redis"
    "github.com/zklapow/bibliophile/models"
    "github.com/pquerna/ffjson/ffjson"
    "fmt"
    "time"
)

const idKey = "id:books"

type Books struct {
    Pool *redis.Pool `inject:""`
}

func (b *Books) Get(id int64) (*models.Book, error) {
    conn := b.Pool.Get()
    return b.getInternal(getKey(id), conn)
}

func (b *Books) GetAndWatch(id int64) (*models.Book, error) {
    conn := b.Pool.Get()

    key := getKey(id)
    if _, err := conn.Do("WATCH", key); err != nil {
        return nil, err
    }

    return b.getInternal(key, conn)
}

func (b *Books) getInternal(key string, conn redis.Conn) (*models.Book, error) {
    data, err := redis.String(conn.Do("GET", key))
    if err != nil {
        if err.Error() == redis.ErrNil.Error() {
            return nil, nil
        }

        return nil, err
    }

    if data == "" {
        return nil, nil
    }

    var book models.Book
    if err := ffjson.Unmarshal([]byte(data), &book); err != nil {
        return nil, err
    }

    return &book, nil
}

func (b *Books) Update(book *models.Book) error {
    conn := b.Pool.Get()
    key := getKey(book.Id)

    data, err := ffjson.Marshal(book)
    if err != nil {
        return err
    }

    conn.Send("MULTI")
    conn.Send("SET", key, data)
    if _, err := conn.Do("EXEC"); err != nil {
        return err
    }

    return nil
}

func (b *Books) Create(book *models.Book) error {
    conn := b.Pool.Get()
    id, err := redis.Int64(conn.Do("INCR", idKey))
    if err != nil {
        return err
    }

    book.Id = id
    if book.Created == 0 {
        book.Created = time.Now().Unix()
    }

    key := getKey(id)
    data, err := ffjson.Marshal(book)
    if err != nil {
        return err
    }

    _, err = conn.Do("SET", key, data)
    if err != nil {
        return err
    }

    return nil
}

func (b *Books) Count() (int64, error) {
    conn := b.Pool.Get()
    count, err := redis.Int64(conn.Do("GET", idKey))
    if err != nil {
        return 0, err
    }

    return count, nil
}

func getKey(id int64) string {
    return fmt.Sprintf("book:%v", id)
}