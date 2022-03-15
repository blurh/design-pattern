package sigleton

import (
    "testing"
)

func TestDesign(t *testing.T) {
    const msg = "hello"
    t.Run("test for sigleton", func(t *testing.T) {
        var got, check *Instance
        go func() { got = GetInstance() }()
        go func() { check = GetInstance() }()
        if got != check {
            t.Error("sigleton test fail")
        }
    })
}
