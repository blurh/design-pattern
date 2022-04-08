package template

type Sender interface {
    Send(string) string
    Get(int) bool
}

type WeChat struct{}

func (w *WeChat) Send(msg string) string {
    return "wechat send:" + " " + msg
}

func (w *WeChat) Get(userid int) bool {
    return true
}

func NewSender() Sender {
    return &WeChat{}
}
