package decorator

import "testing"

func TestDesign(t *testing.T) {
    t.Run("test for decorator", func(t *testing.T) {
        s := NewSender("wechat")
        got := s.Send("hello")
        want := "wechat hello send"
        if got != want {
            t.Error("decorator test fail")
        }
    })
}
