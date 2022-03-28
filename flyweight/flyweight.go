package flyweight

type Pool interface {
    Get(string) any
    Set(string, any) int
}

type PoolMap map[string]any

type ResourcePool struct {
    Resource PoolMap
    Cap      int
}

func (r *ResourcePool) Get(key string) any {
    if value, ok := r.Resource[key]; ok {
        return value
    }
    return ""
}

func (r *ResourcePool) Set(key string, value any) int {
    if _, ok := r.Resource[key]; !ok {
        r.Resource[key] = value
        r.Cap++
        return 0
    } else {
        return 1
    }
}

func NewPool() Pool {
    p := make(PoolMap)
    return &ResourcePool{
        Resource: p,
        Cap:      0,
    }
}
