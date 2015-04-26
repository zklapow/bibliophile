package persist
import (
    "github.com/garyburd/redigo/redis"
    "github.com/zklapow/bibliophile/models"
)

type Books struct {
    Pool *redis.Pool `inject:""`
}

func (b *Books) write(book *models.Book) error {
    return nil
}