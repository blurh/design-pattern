package command

import (
    "reflect"
    "testing"
)

func TestDesign(t *testing.T) {
    t.Run("test for command", func(t *testing.T) {
        q := NewQueue()
        q.Push("start", "pause")
        got := q.Execute()
        want := []string{"do start", "do pause"}
        if !reflect.DeepEqual(want, got) {
            t.Error("command test fail")
        }
    })
}
