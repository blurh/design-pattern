package bridge

type Msg interface {
    Send(string) string
}

type Markdown struct{}

func (m *Markdown) Send(msg string) string {
    s := "markdown" + " " + msg
    return s
}

type Text struct{}

func (t *Text) Send(msg string) string {
    s := "text" + " " + msg
    return s
}

type MsgSender interface {
    SendMsg(string, string) string
}

// 桥接
type WeChat struct {
    Msg Msg
}

func (w *WeChat) SendMsg(sendType, msg string) string {
    switch sendType {
    case "markdown":
        w.Msg = &Markdown{}
    case "text":
        w.Msg = &Text{}
    default:
        return ""
    }
    m := "wechat" + " " + msg
    return w.Msg.Send(m)
}

// 桥接
type DingTalk struct {
    Msg Msg
}

func (d *DingTalk) SendMsg(sendType, msg string) string {
    switch sendType {
    case "markdown":
        d.Msg = &Markdown{}
    case "text":
        d.Msg = &Text{}
    default:
        return ""
    }
    m := "dingtalk" + " " + msg
    return d.Msg.Send(m)
}

func NewSender(sender string) MsgSender {
    switch sender {
    case "wechat":
        return &WeChat{}
    case "dingtalk":
        return &DingTalk{}
    default:
        return nil
    }
}
