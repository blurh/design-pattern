package strategy

type Sender interface {
    Send(msg string) string
}

type Context struct {
    Sender Sender
}

func (c *Context) Send(msg string) string {
    return c.Sender.Send(msg)
}

type Conf struct {
    Url string
}

type WeChat struct {
    *Conf
}

func (w *WeChat) Send(msg string) string {
    // send msg to wechat robot...
    result := "wechat send:" + " " + w.Conf.Url + " " + msg
    return result
}

type DingTalk struct {
    *Conf
}

func (d *DingTalk) Send(msg string) string {
    // send msg to dingding...
    result := "dingding send:" + " " + d.Conf.Url + " " + msg
    return result
}

func NewSender(s Sender) *Context {
    return &Context{
        Sender: s,
    }
}
