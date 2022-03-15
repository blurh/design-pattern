package prototype

type ProtoTypeMsg struct {
    Level string
    Title string
    Msg   string
}

func (m *ProtoTypeMsg) DeepClone() *ProtoTypeMsg {
    return &ProtoTypeMsg{
        Level: m.Level,
        Title: m.Title,
        Msg:   m.Msg,
    }
}

func (m *ProtoTypeMsg) SetLevel(level string) {
    m.Level = level
}

func (m *ProtoTypeMsg) SetTitle(title string) {
    m.Title = title
}

func (m *ProtoTypeMsg) SetMsg(msg string) {
    m.Msg = msg
}

func NewProtoTypeMsg() *ProtoTypeMsg {
    return &ProtoTypeMsg{}
}
