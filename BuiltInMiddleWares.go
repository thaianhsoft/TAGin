package TAGin

import (
	"log"
	"time"
)

func LogHandler() HandlerFunc {
	handler := func(u *UserContext) {
		t := time.Now()
		u.NextHandler()
		log.Printf("[%v]:[%v] finished [%v]", u.Req.Method, u.Req.URL.Path, time.Since(t))
	}
	return handler
}