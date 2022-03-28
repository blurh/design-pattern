package flyweight

import "testing"

func TestDesign(t *testing.T) {
    t.Run("test for flyweight", func(t *testing.T) {
        p := NewPool()
        p.Set("k", "v")
        p.Set("kx", "vx")
        got := p.Get("kx")
        want := "vx"
        if got.(string) != want {
            t.Error("flyweight test fail")
        }
    })
}
