package observer

import "sync"

type DataEvent struct {
    Topic string
    Data  any
}

type DataChan chan DataEvent

type EventBus struct {
    SubScribers sync.Map
}

func (eb *EventBus) SubScribe(topic string, ch DataChan) {
    if _, found := eb.SubScribers.Load(topic); !found {
        ch := append([]DataChan{}, ch)
        eb.SubScribers.Store(topic, ch)
    }
}

func (eb *EventBus) Publish(topic string, data any) {
    if chanSlice, found := eb.SubScribers.Load(topic); found {
        go func(data DataEvent, ch []DataChan) {
            for _, ch := range chanSlice.([]DataChan) {
                ch <- data
            }
        }(DataEvent{Topic: topic, Data: data}, chanSlice.([]DataChan))
    }
}

func GetDataEvent(ch DataChan) DataEvent {
    return <-ch
}

func NewDataChan() DataChan {
    return make(DataChan)
}

func NewEventBus() *EventBus {
    return &EventBus{
        SubScribers: sync.Map{},
    }
}
