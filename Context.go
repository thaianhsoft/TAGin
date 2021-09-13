package TAGin

import "net/http"
type HandlerFunc func(*UserContext)

type UserContext struct {
	Req *http.Request
	Writer *http.ResponseWriter
	Params map[string]interface{}
}
