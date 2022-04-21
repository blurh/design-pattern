package visitor

import "testing"

func TestDesign(t *testing.T) {
    t.Run("test for visitor", func(t *testing.T) {
        compressor := NewCompressor()
        reader := NewReader("xx.csv")
        got := reader.Visit(compressor)
        want := "csv"
        if got != want {
            t.Error("visitor test fail")
        }
    })
}
