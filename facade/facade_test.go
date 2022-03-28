package facade

import "testing"

func TestDesign(t *testing.T) {
    sender := NewSender("wechat")
    got := sender.LoginAndSend("username", "password", "hello")
    want := "wechat hello"
    if got != want {
        t.Errorf("facade test fail")
    }
}
