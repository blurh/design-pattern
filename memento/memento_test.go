package memento

import (
    "reflect"
    "testing"
)

func TestDesign(t *testing.T) {
    t.Run("test for memento", func(t *testing.T) {
        data := NewData()
        data.Add("x")
        data.Add("y")
        s := data.Snapshot()
        data.Add("z")
        data.Restore(s)
        got := data.Content
        want := []string{"x", "y"}
        if !reflect.DeepEqual(got, want) {
            t.Error("memento test fail")
        }
    })
}
