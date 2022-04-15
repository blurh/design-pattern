package chainresponsibility

type HandleFunc func()

type RouterGroup struct {
    Handles    []HandleFunc
    ChainCount int
}

func (r *RouterGroup) Use(f ...HandleFunc) {
    r.Handles = append(r.Handles, f...)
    r.ChainCount += len(f)
}

func (r *RouterGroup) Next() {
    if r.ChainCount == 0 {
        return
    }
    r.Handles[0]()
    r.ChainCount--
    r.Handles = append(r.Handles[:0], r.Handles[1:]...)
}

func NewRouterGroup() *RouterGroup {
    return &RouterGroup{
        ChainCount: 0,
    }
}
