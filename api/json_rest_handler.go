package api
import (
    "net/http"
    "github.com/zklapow/bibliophile/util"
    "github.com/pquerna/ffjson/ffjson"
    "github.com/Sirupsen/logrus"
)

type JsonRestHandler interface {
    Get(resp http.ResponseWriter, req *http.Request) (int, interface{})
    Post(resp http.ResponseWriter, req *http.Request) (int, interface{})
    Logger() (*logrus.Logger)
}

type JsonRestHttpHandler struct {
    Handler JsonRestHandler
    Log     *logrus.Logger `inject:""`
}

func (h *JsonRestHttpHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
    code := 0
    var val interface{}

    switch req.Method {
        case "GET":
            code, val = h.Handler.Get(resp, req)
        case "POST":
            code, val = h.Handler.Post(resp, req)
        default:
            code = http.StatusMethodNotAllowed
    }

    log := h.Handler.Logger().WithField("method", req.Method)

    var err error
    var data []byte
    if val != nil {
        util.ReturnsJSON(resp)
        data, err = ffjson.Marshal(val)
        if err != nil {
            log.Errorf("Error marshalling JSON", err)
            resp.WriteHeader(http.StatusInternalServerError)
            return
        }
    }

    if code != 0 {
        resp.WriteHeader(code)
    }
    if data != nil {
        resp.Write(data)
    }
}

func (h *JsonRestHttpHandler) Get(resp http.ResponseWriter, req *http.Request) (int, interface{}) {
    return http.StatusMethodNotAllowed, nil
}

func (h *JsonRestHttpHandler) Post(resp http.ResponseWriter, req *http.Request) (int, interface{}) {
    return http.StatusMethodNotAllowed, nil
}

func (h *JsonRestHttpHandler) Logger() *logrus.Logger {
    return h.Log
}