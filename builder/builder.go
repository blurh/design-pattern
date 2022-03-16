package builder

import "errors"

type Msg struct {
    maxLength int
    content   string
    toUser    int
}

type Sender interface {
    Send() string
}

type SetMsg interface {
    setMaxLength(length int) SetMsg
    setContent(msg string) SetMsg
    setUser(userid int) SetMsg
    getErr() error
    Build() Sender
}

type WeChatMsg struct {
    Msg Msg
}

func (w *WeChatMsg) Send() string {
    return w.Msg.content
}

type SetWeChatMsg struct {
    Msg Msg
    err error
}

func (s *SetWeChatMsg) setMaxLength(length int) SetMsg {
    s.Msg.maxLength = length
    return s
}

func (s *SetWeChatMsg) setContent(msg string) SetMsg {
    s.Msg.content = msg
    return s
}

func (s *SetWeChatMsg) setUser(userid int) SetMsg {
    s.Msg.toUser = userid
    return s
}

func (s *SetWeChatMsg) getErr() error {
    return s.err
}

func (s *SetWeChatMsg) Build() Sender {
    contentLen := len(s.Msg.content)
    if contentLen > s.Msg.maxLength {
        s.err = errors.New("long")
        return nil
    }
    return &WeChatMsg{
        Msg: Msg{
            maxLength: s.Msg.maxLength,
            content:   s.Msg.content,
            toUser:    s.Msg.toUser,
        },
    }
}

type MsgDirector struct{}

func (d *MsgDirector) MsgConstruct(builder *SetWeChatMsg) (Sender, error) {
    b := builder.
        setContent("wechat build send").
        setMaxLength(30).
        setUser(1)
    if b == nil && b.getErr() != nil {
        return nil, b.getErr()
    }
    return b.Build(), nil
}
