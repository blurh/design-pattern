package proxy

type Sender interface {
    Send(string) string
}

type WeChat struct{}

func (w *WeChat) Send(msg string) string {
    r := "wechat send: " + msg
    return r
}

type DingTalk struct{}

func (d *DingTalk) Send(msg string) string {
    r := "dingtalk send: " + msg
    return r
}

type Proxy struct {
    Sender Sender
}

func (p *Proxy) checkPermission() bool {
    return true
}

func (p *Proxy) Send(msg string) string {
    // do something
    if check := p.checkPermission(); !check {
        return ""
    }
    r := p.Sender.Send(msg)
    return r
}

func NewProxySender(sender Sender) *Proxy {
    return &Proxy{
        Sender: sender,
    }
}
