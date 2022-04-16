package state

type Status interface {
    SendMsg() string
    SendSms() string
    TopUP(int) bool
}

// 状态机
type Account struct {
    Balance int
    Status  Status
}

func (a *Account) CheckBalance() bool {
    if a.Balance <= 0 {
        return false
    }
    return true
}

func (a *Account) SetAccState() {
    if a.CheckBalance() {
        a.Status = &Nornal{Account: *a}
    } else {
        a.Status = &Disable{Account: *a}
    }
}

func (a *Account) SendMsg() string {
    return a.Status.SendMsg()
}

func (a *Account) SendSms() string {
    return a.Status.SendSms()
}

func (a *Account) TopUP(amount int) bool {
    return a.Status.TopUP(amount)
}

type Nornal struct {
    Account Account
}

func (n *Nornal) SendMsg() string {
    if !n.Account.CheckBalance() {
        n.Account.SetAccState()
        return "balance not enough, send msg fail"
    }
    n.Account.Balance--
    return "send msg success"
}

func (n *Nornal) SendSms() string {
    if !n.Account.CheckBalance() {
        n.Account.SetAccState()
        return "balance not enough, send sms fail"
    }
    n.Account.Balance--
    return "send sms success"
}

func (n *Nornal) TopUP(amount int) bool {
    n.Account.Balance += amount
    return true
}

type Disable struct {
    Account Account
}

func (d *Disable) SendMsg() string {
    return "account disabled, send msg fail"
}

func (d *Disable) SendSms() string {
    return "account disabled, send sms fail"
}

func (d *Disable) TopUP(amount int) bool {
    d.Account.Balance += amount
    return true
}

func NewAccount(amount int) *Account {
    a := &Account{
        Balance: amount,
    }
    a.SetAccState()
    return a
}
