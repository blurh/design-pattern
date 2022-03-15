package prototype

import (
    "testing"
)

func TestDesign(t *testing.T) {
    t.Run("test for proto type", func(t *testing.T) {
        pMsg := NewProtoTypeMsg()
        pMsg.SetLevel("warning")
        pMsg.SetTitle("title")
        pMsg.SetMsg("msg")
        pMsgClone := pMsg.DeepClone()
        pMsgClone.SetLevel("error")
        if pMsgClone.Level != "error" || pMsg.Title != pMsgClone.Title || pMsg.Msg != pMsgClone.Msg {
            t.Error("proto type test fail")
        }
    })
}
