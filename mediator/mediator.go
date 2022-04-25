package mediator

const URL = "http://xxx"

type TextFiled struct {
    Value string
}

func (i *TextFiled) Input(v string) {
    i.Value = v
}

type Selection struct {
    Values string
}

func (s *Selection) Checked(checked string) {
    s.Values = checked
}

type Dialog struct {
    Account        *TextFiled
    Passwd         *TextFiled
    RepeatPassword *TextFiled
    Type           *Selection
}

// 中介
func (d *Dialog) Submit() string {
    checked := d.Type.Values
    switch checked {
    case "login":
        return URL + "?acc=" + d.Account.Value + "&passwd=" + d.Passwd.Value
    case "register":
        return URL + "?acc=" + d.Account.Value + "&passwd=" + d.Passwd.Value + "&repeat=" + d.RepeatPassword.Value
    default:
        return ""
    }
}

func NewTextField() *TextFiled {
    return &TextFiled{}
}

func NewSelection() *Selection {
    return &Selection{}
}

func NewDialog() *Dialog {
    return &Dialog{}
}
