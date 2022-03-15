package abstract_factory

type User interface {
    Get(userid int) bool
    Send(msg string) string
}

type WeChatUser struct{}

func (w *WeChatUser) Get(userid int) bool {
    return true
}

func (w *WeChatUser) Send(msg string) string {
    return "wechat user send"
}

type DingTalkUser struct{}

func (d *DingTalkUser) Get(userid int) bool {
    return true
}

func (d *DingTalkUser) Send(msg string) string {
    return "dingtalk user send"
}

type Group interface {
    List(groupid int) []string
    Send(msg string) string
}

type WeChatGroup struct{}

func (w *WeChatGroup) List(groupid int) []string {
    return []string{"wechat", "group"}
}

func (w *WeChatGroup) Send(msg string) string {
    return "wechat group send"
}

type DingTalkGroup struct{}

func (d *DingTalkGroup) List(groupid int) []string {
    return []string{"dingtalk", "group"}
}

func (d *DingTalkGroup) Send(msg string) string {
    return "dingtalk group send"
}

type Chat interface {
    CreateUser() User
    CreateGroup() Group
}

type WeChatCreater struct{}

func (c *WeChatCreater) CreateUser() User {
    return &WeChatUser{}
}

func (c *WeChatCreater) CreateGroup() Group {
    return &WeChatGroup{}
}

type DingTalkCreater struct{}

func (c *DingTalkCreater) CreateUser() User {
    return &DingTalkUser{}
}

func (c *DingTalkCreater) CreateGroup() Group {
    return &DingTalkGroup{}
}

func NewChat(chat string) Chat {
    switch chat {
    case "wechat":
        return &WeChatCreater{}
    case "dingtalk":
        return &DingTalkCreater{}
    default:
        return nil
    }
}
