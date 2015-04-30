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

type InjectLogger struct {
    Log *logrus.Logger `inject:""`
}

func (i InjectLogger) Logger() *logrus.Logger {
    return i.Log
}

type GetNotAllowed struct {
}

func (g GetNotAllowed) Get(resp http.ResponseWriter, req *http.Request) (int, interface{}) {
    return http.StatusMethodNotAllowed, nil
}

type PostNotAllowed struct {
}

func (p PostNotAllowed) Post(resp http.ResponseWriter, req *http.Request) (int, interface{}) {
    return http.StatusMethodNotAllowed, nil
}

func requestHandler(h JsonRestHandler) http.HandlerFunc {
    return func (resp http.ResponseWriter, req *http.Request) {
        code := 0
        var val interface{}

        switch req.Method {
            case "GET":
                code, val = h.Get(resp, req)
            case "POST":
                code, val = h.Post(resp, req)
            default:
                code = http.StatusMethodNotAllowed
        }

        log := h.Logger().WithField("method", req.Method)

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
}
