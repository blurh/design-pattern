package simplefactory

import (
    "testing"
)

func TestDesign(t *testing.T) {
    const msg = "hello"
    t.Run("test for simple factory", func(t *testing.T) {
        sender := NewSender("wechat")
        got := sender.Send(msg)
        want := "wechat send:" + " " + msg
        if got != want {
            t.Error("simple factory test fail")
        }
    })
}
