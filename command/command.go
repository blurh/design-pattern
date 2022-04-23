package command

import "time"

type Command interface {
    Execute() string
}

type Start struct{}

func (s *Start) Execute() string {
    return "do start"
}

type Pause struct{}

func (p *Pause) Execute() string {
    return "do pause"
}

func NewCommand(comm string) Command {
    switch comm {
    case "start":
        return &Start{}
    case "pause":
        return &Pause{}
    default:
        return nil
    }
}

type Queue chan string

func NewQueue() Queue {
    return make(Queue, 10)
}

func (q Queue) Push(commands ...string) {
    go func() {
        for _, comm := range commands {
            q <- comm
        }
    }()
}

func (q Queue) Execute() (result []string) {
    for {
        select {
        case c := <-q:
            r := NewCommand(c).Execute()
            result = append(result, r)
        case <-time.After(time.Second * 1 / 5):
            return
        }
    }
}
