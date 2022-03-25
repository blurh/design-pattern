package adapter

import "testing"

func TestDesign(t *testing.T) {
    t.Run("test for adapter", func(t *testing.T) {
        s := NewSender("wechat")
        got := s.Send("hello")
        want := "wechat robot hello"
        if got != want {
            t.Error("adapter test fail")
        }
    })
}
