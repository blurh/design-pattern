package proxy

import "testing"

func TestDesign(t *testing.T) {
    t.Run("test for proxy", func(t *testing.T) {
        msg := "hello"
        chat := &WeChat{}
        sender := NewProxySender(chat)
        got := sender.Send(msg)
        want := "wechat send: " + msg
        if got != want {
            t.Error("proxy test fail")
        }
    })
}
