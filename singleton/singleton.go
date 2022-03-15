package sigleton

import "sync"

type Instance struct{}

var (
    once     sync.Once
    instance *Instance
)

func GetInstance() *Instance {
    once.Do(func() {
        instance = new(Instance)
    })
    return instance
}
