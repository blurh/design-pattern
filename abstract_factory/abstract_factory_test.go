package abstract_factory

import (
    "testing"
)

func TestDesign(t *testing.T) {
    const msg = "hello"
    t.Run("test for abstract factory", func(t *testing.T) {
        factory := NewChat("wechat")
        user := factory.CreateUser()
        got := user.Send(msg)
        want := "wechat user send"
        if got != want {
            t.Error("abstract factory test fail")
        }
    })
}
