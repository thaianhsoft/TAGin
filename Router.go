package TAGin

import (
	"github.com/thaianhsoft/tads/graph/directed/prefixtrie"
	"log"
	"net/http"
	"strings"
)

type router struct {
	tries   map[string]*prefixtrie.Trie
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		tries:   make(map[string]*prefixtrie.Trie),
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
	trie, _ := r.tries[method]
	if trie == nil {
		trie = prefixtrie.NewTrie()
		r.tries[method] = trie
	}
	parts := r.parse(pattern)
	trie.InsertPath(parts, pattern)
	log.Println(key)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, pattern string) (*string, map[string]string) {
	if r.tries[method] == nil {
		return nil, nil
	}
	parts := r.parse(pattern)
	trie := r.tries[method]
	if paths := trie.SearchPath(parts); paths != nil {
		parsePath := r.parse(*paths)
		params := make(map[string]string)
		for i, path := range parsePath {
			if path[0] == ':' {
				params[path[1:]] = parts[i]
			}
			if path[0] == '$' {
				params[path[1:]] = strings.Join(parts[i:], "/")
			}
		}
		return paths, params
	}
	return nil, nil
}

func (r *router) handle(c *UserContext) {
	user_method := c.Req.Method
	user_path := c.Req.URL.Path
	if path, params := r.getRoute(user_method, user_path); path != nil && params != nil {
		key := user_method + "-" + (*path)
		log.Println(params)
		c.Params = params
		handler := r.handlers[key]
		c.middlewares = append(c.middlewares, handler)
	} else {
		c.middlewares = append(c.middlewares, func(u *UserContext) {
			http.Error(u.Writer, "404 not found", 404)
		})
	}
	c.NextHandler()
}
