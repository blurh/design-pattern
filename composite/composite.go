package composite

type FileSystem interface {
    Add(FileSystem)
    List() []string
}

type File struct {
    Name string
}

func (f *File) Add(fs FileSystem) {}

func (f *File) List() []string {
    return []string{f.Name}
}

type Dictory struct {
    Name string
    Node []FileSystem
}

func (d *Dictory) Add(node FileSystem) {
    d.Node = append(d.Node, node)
}

func (d *Dictory) List() []string {
    ls := []string{d.Name}
    for _, node := range d.Node {
        ls = append(ls, node.List()...)
    }
    return ls
}

func NewFileSystem() FileSystem {
    return &Dictory{Name: "root"}
}

func NewDir(name string) FileSystem {
    return &Dictory{Name: name}
}

func NewFile(name string) FileSystem {
    return &File{Name: name}
}
