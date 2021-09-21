package TAGin

import (
	"log"
	"testing"
)

func TestRouter(t *testing.T) {
	router := newRouter()
	router.addRoute("GET", "/api/view/:id/get/:name", func(context *UserContext) {

	})
	router.addRoute("GET", "api/view/:id/post/:name", func(context *UserContext) {

	})
	router.addRoute("GET", "api/view/$file", func(context *UserContext) {

	})
	{
		_, params := router.getRoute("GET", "/api/view/3/get/thaianh")
		log.Println(params)
		// PASSED
	}
	{
		_, params := router.getRoute("GET", "/api/view/3/post/minhcanh")
		log.Println(params)
		// PASSED
	}
	{
		ok, params := router.getRoute("GET", "/api/show/3")
		log.Println(params, ok)
	}
	{
		ok, params := router.getRoute("GET", "/api/view/css")
		log.Println(params, ok)
	}
}
