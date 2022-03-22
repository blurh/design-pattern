package bridge

import "testing"

func TestDesign(t *testing.T) {
    t.Run("test for bridge", func(t *testing.T) {
        sender := NewSender("wechat")
        got := sender.SendMsg("markdown", "hello")
        want := "markdown wechat hello"
        if got != want {
            t.Error("test bridge fail")
        }
    })
}
