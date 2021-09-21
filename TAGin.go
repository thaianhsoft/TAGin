package TAGin

import (
	"log"
	"net/http"
	path2 "path"
	"strings"
)

type group struct {
	*engine
	parent *group
	middlewares []HandlerFunc
	prefix string
}

func (g *group) Group(prefix string) *group {
	e := g.engine
	new_group := &group{
		engine:      e,
		parent:      g,
		prefix:      g.prefix+prefix,
	}
	e.groups = append(e.groups, new_group)
	return new_group
}
func (g *group) Static(path string, folder string) {
	handler := g.staticHandler(path, http.Dir(folder))
	urlPattern := path2.Join(path, "/$file")
	g.GET(urlPattern, handler)
}
func (g *group) staticHandler(pattern string, fs http.FileSystem) HandlerFunc {
	p := path2.Join(g.prefix, pattern)
	fileServer := http.StripPrefix(p, http.FileServer(fs))
	handler := func (u *UserContext) {
		file := u.Param("file")
		if _, err := fs.Open(file); err != nil {
			log.Println(err)
		}
		log.Printf("fileServer address %v\n", fileServer)
		fileServer.ServeHTTP(u.Writer, u.Req)
	}
	return handler
}

func (g *group) Use(handlers ...HandlerFunc) {
	g.middlewares = append(g.middlewares, handlers...)
}
func (g *group) GET(pattern string, handler HandlerFunc) {
	g.addRoute(http.MethodGet, pattern, handler)
}

func (g *group) POST(pattern string, handler HandlerFunc) {
	g.addRoute(http.MethodPost, pattern, handler)
}

func (g *group) addRoute(method string, pattern string, handler HandlerFunc) {
	prefixPattern := g.prefix+pattern
	g.router.addRoute(method, prefixPattern, handler)
}
type engine struct {
	router *router
	*group
	groups []*group
}

func NewTAGin() *engine {
	engine := &engine{router: newRouter()}
	engine.group = &group{engine: engine}
	engine.groups = make([]*group, 0)
	engine.groups = append(engine.groups, engine.group)
	return engine
}

func (e *engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userContext := newUserContext(w, r)
	middlewares := make([]HandlerFunc, 0)
	for _, group := range e.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	userContext.middlewares = middlewares
	log.Println(userContext.index)
	e.router.handle(userContext)
}

func (e *engine) Run(addr string) {
	http.ListenAndServe(addr, e)
}

