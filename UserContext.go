package TAGin

import (
	"encoding/json"
	"log"
	"net/http"
)

type HandlerFunc func(*UserContext)
type Message map[string]interface{}
type UserContext struct {
	Req        *http.Request
	Writer     http.ResponseWriter
	Params     map[string]string
	middlewares []HandlerFunc
	index int
}
func newUserContext(w http.ResponseWriter, r *http.Request) *UserContext {
	return &UserContext{
		Req:    r,
		Writer: w,
		index: -1,
	}
}
func (u *UserContext) Status(code int) {
	u.Writer.WriteHeader(code)
}

func (u *UserContext) SetHeader(key string, value string) {
	u.Writer.Header().Set(key, value)
}

func (u *UserContext) HTML(code int, content []byte) {
	u.Status(code)
	u.SetHeader("Content-Type", "text/html")
	u.Writer.Write(content)
}

func (u *UserContext) JSON(code int, message Message) {
	u.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(u.Writer)
	if err := encoder.Encode(message); err != nil {
		http.Error(u.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (u *UserContext) Param(key string) string {
	if v, ok := u.Params[key]; ok {
		return v
	} else {
		return v
	}
}

func (u *UserContext) NextHandler() {
	u.index++
	for ; u.index < len(u.middlewares); u.index++ {
		log.Printf("index %v and len %v\n", u.index, len(u.middlewares))
		u.middlewares[u.index](u)
	}
}

