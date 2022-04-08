package template

import "testing"

func TestDesign(t *testing.T) {
    t.Run("test for template", func(t *testing.T) {
        sender := NewSender()
        got := sender.Send("hello")
        want := "wechat send: hello"
        if got != want {
            t.Error("template test fail")
        }
    })
}
