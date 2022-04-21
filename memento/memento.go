package memento

type Data struct {
    Content []string
}

type Snapshot struct {
    Content []string
}

func (d *Data) Add(data string) {
    d.Content = append(d.Content, data)
}

func (d *Data) Snapshot() *Snapshot {
    return NewSnapshot(d.Content)
}

func (d *Data) Restore(s *Snapshot) {
    d.Content = append([]string{}, s.Content...)
}

func NewData() *Data {
    return &Data{}
}

func NewSnapshot(datas []string) *Snapshot {
    return &Snapshot{Content: datas}
}
