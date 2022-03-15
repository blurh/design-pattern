package strategy

import (
    "testing"
)

func TestDesign(t *testing.T) {
    const msg = "hello"
    t.Run("test for strategy", func(t *testing.T) {
        wechatRobot := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xx"
        sender := NewSender(&WeChat{
            Conf: &Conf{
                Url: wechatRobot,
            },
        })
        got := sender.Send(msg)
        want := "wechat send: " + wechatRobot + " " + msg
        if got != want {
            t.Error("strategy test fail")
        }
    })
}
