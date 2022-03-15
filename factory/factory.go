package factory

type Sender interface {
    Send(msg string) string
}

type Msg interface {
    CreateSender() Sender
}

type Conf struct {
    Url string
}

type WeChatSender struct {
    Conf
}

func (w *WeChatSender) Send(msg string) string {
    result := "wechat send:" + " " + msg
    return result
}

type CreateWeChatSender struct{}

func (c *CreateWeChatSender) CreateSender() Sender {
    return &WeChatSender{
        Conf{
            Url: "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xx",
        },
    }
}

type DingTalkSender struct {
    Conf
}

func (d *DingTalkSender) Send(msg string) string {
    result := "dingding send:" + " " + msg
    return result
}

type CreateDingTalkSender struct{}

func (c *CreateDingTalkSender) CreateSender() Sender {
    return &DingTalkSender{
        Conf{
            Url: "https://oapi.dingtalk.com/robot/send?access_token=xx",
        },
    }
}

func NewSender(webhook string) Msg {
    switch webhook {
    case "wechat":
        return &CreateWeChatSender{}
    case "dingding":
        return &CreateDingTalkSender{}
    default:
        return nil
    }
}
