package observer

import "testing"

func TestDesign(t *testing.T) {
    t.Run("test for observer", func(t *testing.T) {
        topic := "observerTopic"
        data := "xxx"

        eb := NewEventBus()
        ch := NewDataChan()

        eb.SubScribe(topic, ch)
        go eb.Publish(topic, data)

        got := GetDataEvent(ch)
        want := data
        if got.Topic != topic || got.Data.(string) != want {
            t.Error("observer test fail")
        }
    })
}
