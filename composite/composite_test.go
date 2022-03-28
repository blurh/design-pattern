package composite

import (
    "reflect"
    "testing"
)

func TestDesign(t *testing.T) {
    t.Run("test for composite", func(t *testing.T) {
        fs := NewFileSystem()
        file := NewFile("file")
        dir := NewDir("dir")
        leaf := NewDir("leaf")
        anotherFile := NewFile("another file")
        fs.Add(file)
        fs.Add(dir)
        fs.Add(leaf)
        leaf.Add(anotherFile)
        got := fs.List()
        want := []string{"root", "file", "dir", "leaf", "another file"}
        if !reflect.DeepEqual(got, want) {
            t.Error("composite test fail")
        }
    })
}
