package adapter

type Sender interface {
    Send(string) string
}

type WeChatRobot struct{}

func (wr *WeChatRobot) Robot(msg string) string {
    m := "robot" + " " + msg
    return m
}

// 适配
type WeChat struct {
    Sender WeChatRobot
}

func (w *WeChat) Send(msg string) string {
    m := w.Sender.Robot(msg)
    return "wechat" + " " + m
}

type DingTalkRobot struct{}

func (dr *DingTalkRobot) WebHook(msg string) string {
    m := "webhook" + " " + msg
    return m
}

// 适配
type DingTalk struct {
    Sender DingTalkRobot
}

func (d *DingTalk) Send(msg string) string {
    m := d.Sender.WebHook(msg)
    return "dingtalk" + " " + m
}

func NewSender(chat string) Sender {
    switch chat {
    case "wechat":
        return &WeChat{}
    case "dingtalk":
        return &DingTalk{}
    default:
        return nil
    }
}
