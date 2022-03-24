package decorator

type Msg interface {
    Send(string) string
}

type Sender struct{}

func (s *Sender) Send(msg string) string {
    m := msg + " " + "send"
    return m
}

type WeChat struct {
    Msg
}

// 装饰
func (w *WeChat) Send(msg string) string {
    w.Msg = &Sender{}
    m := "wechat" + " " + msg
    return w.Msg.Send(m)
}

type DingTalk struct {
    Msg
}

// 装饰
func (d *DingTalk) Send(msg string) string {
    d.Msg = &Sender{}
    m := "dingtalk" + " " + msg
    return d.Msg.Send(m)
}

func NewSender(chat string) Msg {
    switch chat {
    case "wechat":
        return &WeChat{}
    case "dingtalk":
        return &DingTalk{}
    default:
        return nil
    }
}
