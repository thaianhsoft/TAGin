package TAGin

type graph struct {
	adj map[interface{}][]interface{}
}

func newGraph() *graph {
	return &graph{
		adj: make(map[interface{}][]interface{}, 0),
	}
}

func (g *graph) addEdge(from interface{}, to interface{}) {
	if _, ok := g.adj[from]; !ok {
		g.adj[from] = make([]interface{}, 0)
		g.adj[from] = append(g.adj[from], to)
		return
	}
	return
}

func (g *graph) addEmptyEdge(source interface{}) {
	if _, ok := g.adj[source]; !ok {
		g.adj[source] = make([]interface{}, 0)
		return
	}
	return
}

func (g *graph) insert(parts interface{}) {
	castedParts, can := parts.([]string)
	if !can {
		return
	}
	_, exist := g.adj[castedParts[0]]
	if !exist {
		g.adj[castedParts[0]] = make([]interface{}, 0)
	}
	s := 0
	t := 1
	for t < len(castedParts) {
		if g.adj[castedParts[t]] == nil {
			g.adj[castedParts[t]] = make([]interface{}, 0)
			g.adj[castedParts[s]] = append(g.adj[castedParts[s]], castedParts[t])
		}
		s = t
		t++
	}
}

func (g *graph) search(parts interface{}) []string {
	castedParts, can := parts.([]string)
	if !can {
		return nil
	}
	if _, ok := g.adj[castedParts[0]]; !ok {
		return nil
	}

	st := newStack()
	st.push(castedParts[0])
	index := 1
	list := []string{castedParts[0]}
	for !st.empty() && index < len(castedParts) {
		top := st.pop()
		if g.adj[top] != nil {
			for _, u := range g.adj[top] {
				if u == castedParts[index] || u.(string)[0] == ':' || u.(string)[0] == '*' {
					st.push(u)
					list = append(list, u.(string))
				}
			}
			index++
		}
	}
	if len(list) != len(castedParts) {
		return nil
	}
	return list
}

type element struct {
	v    interface{}
	down *element
}

type stack struct {
	top *element
	l   int
}

func newStack() *stack {
	return &stack{
		l: 0,
	}
}

func (s *stack) push(v interface{}) {
	if s.l == 0 {
		s.top = &element{v: v}
		s.l++
		return
	}
	new_top := &element{v: v}
	down := s.top
	new_top.down = down
	s.top = new_top
	s.l++
	return
}

func (s *stack) pop() interface{} {
	down := s.top.down
	pop_top := s.top
	s.top = down
	s.l--
	return pop_top.v
}

func (s *stack) empty() bool {
	return s.l == 0
}
