package builder

import (
    "testing"
)

func TestDesign(t *testing.T) {
    t.Run("test for builder", func(t *testing.T) {
        b := &SetWeChatMsg{}
        d := &MsgDirector{}
        wechat, err := d.MsgConstruct(b)
        if err != nil {
            t.Error("check builder fail")
            return
        }
        got := wechat.Send()
        want := "wechat build send"
        if got != want {
            t.Error("builder test fail")
        }
    })
}
