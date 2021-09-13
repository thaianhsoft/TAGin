package TAGin

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	graphs   map[string]*graph // one get or post is one graph
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		graphs:   make(map[string]*graph),
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *router) parse(pattern string) []string {
	parts := make([]string, 0)
	patternParts := strings.Split(pattern, "/")
	for _, part := range patternParts {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	graph, _ := r.graphs[method]
	if graph == nil {
		graph = newGraph()
		r.graphs[method] = graph
	}
	parts := r.parse(pattern)
	graph.insert(parts)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, pattern string) (bool, map[string]interface{}) {
	if r.graphs[method] == nil {
		return false, nil
	}
	parts := r.parse(pattern)
	graph := r.graphs[method]
	if list := graph.search(parts); list != nil {
		params := make(map[string]interface{})
		for i, v := range list {
			if v[0] == ':' {
				params[v] = parts[i]
			}
		}
		return true, params
	}
	return false, nil
}

func (r *router) handle(c *UserContext) {
	method := c.Req.Method
	pattern := c.Req.URL.Path
	if ok, params := r.getRoute(method, pattern); ok {
		c.Params = params
		r.handlers[method+"-"+pattern](c)
		log.Println(params)
	} else {
		http.Error(c.Writer, "404 not found", 404)
	}
}
