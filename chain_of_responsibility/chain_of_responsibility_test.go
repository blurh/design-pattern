package chainresponsibility

import "testing"

func TestDesign(t *testing.T) {
    t.Run("test for chain of responsibity", func(t *testing.T) {
        ch := make(chan string, 3)
        f1 := func() {
            ch <- "f1"
        }
        f2 := func() {
            ch <- "f2"
        }
        f3 := func() {
            ch <- "f3"
        }

        r := NewRouterGroup()
        r.Use(f1, f2)
        r.Use(f3)
        for i := 0; i <= 2; i++ {
            r.Next()
        }

        var got string
        for i := range ch {
            got += i
            if len(ch) == 0 {
                break
            }
        }

        want := "f1f2f3"
        if got != want {
            t.Error("chain of responsibity test fail")
        }
    })
}
