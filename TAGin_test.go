package TAGin

import (
	"log"
	"testing"
)

func TestServer(t *testing.T) {
	TAGin := NewTAGin()
	group := TAGin.Group("/api")
	group.Use(LogHandler())
	{
		group.GET("/view/:id/get/:name", func(c *UserContext) {
			c.JSON(200, Message{
				"id": c.Param("id"),
				"name": c.Param("name"),
			})
			log.Println("user handler finished")
		})
		group.GET("/view/thaianh", func(c *UserContext) {
			c.HTML(200, []byte(`
				<div>
					<h1>Xin CHAO</h1>
				<div>
			`))
		})
		group.POST("/:id", func(c *UserContext) {
			c.JSON(200, Message{
				"id": c.Param("id"),
			})
		})
	}
	v2 := TAGin.Group("/view")
	v2.Use(LogHandler())
	{
		v2.Static("/asset/", "./thaianh/")
	}
	TAGin.Run("localhost:8080")
}