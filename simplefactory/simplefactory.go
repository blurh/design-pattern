package simplefactory

type Sender interface {
    Send(msg string) string
}

type DingTalk struct{}

func (d *DingTalk) Send(msg string) string {
    // send msg to dingding...
    result := "dingding send:" + " " + msg
    return result
}

type WeChat struct{}

func (w *WeChat) Send(msg string) string {
    // send msg to wechat robot...
    result := "wechat send:" + " " + msg
    return result
}

func NewSender(webhook string) Sender {
    switch webhook {
    case "wechat":
        return new(WeChat)
    case "dingding":
        return new(DingTalk)
    default:
        panic("unknow webhook type")
    }
}
