package facade

type Sender interface {
    Login(username, passowrd string) bool
    Send(string) string
    LoginAndSend(string, string, string) string
}

type WeChat struct{}

func (w *WeChat) Login(username, passowrd string) bool {
    return true
}

func (w *WeChat) Send(msg string) string {
    return "wechat" + " " + msg
}

// 门面
func (w *WeChat) LoginAndSend(username, passowrd, msg string) string {
    if w.Login(username, passowrd) {
        return w.Send(msg)
    }
    return ""
}

func NewSender(chat string) Sender {
    switch chat {
    case "wechat":
        return &WeChat{}
    default:
        return nil
    }
}
