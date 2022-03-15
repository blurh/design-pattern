package factory

import (
    "testing"
)

func TestDesign(t *testing.T) {
    const msg = "hello"
    t.Run("test for factory", func(t *testing.T) {
        sender := NewSender("wechat")
        send := sender.CreateSender()
        got := send.Send(msg)
        want := "wechat send:" + " " + msg
        if got != want {
            t.Error("factory test fail")
        }
    })
}
