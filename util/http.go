package util
import "net/http"

func ReturnsJSON(resp http.ResponseWriter) {
    resp.Header().Add("Content-Type", "application/json")
}
