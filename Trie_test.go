package TAGin

import (
	"log"
	"testing"
)

func TestInsert(t *testing.T) {
	router := newRouter()
	router.addRoute("GET", "/api/view/:name", func(c *UserContext) {
		log.Println("hello")
	})
	router.addRoute("GET", "/api/view/sukvat", func(c *UserContext) {})

	router.getRoute("GET", "/api/view/thaianh/hihi")
}
