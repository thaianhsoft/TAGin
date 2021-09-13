package TAGin

type engine struct {
	router *router
}

func NewTAGin() *engine {
	return &engine{
		router: newRouter(),
	}
}
