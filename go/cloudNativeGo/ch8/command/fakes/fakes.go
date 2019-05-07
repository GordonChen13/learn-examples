package fakes

type FakeQueueDispatcher struct {
	Messages []interface{}
}

func NewFakeQueueDispatcher() (dispatcher *FakeQueueDispatcher) {
	dispatcher = &FakeQueueDispatcher{}
	dispatcher.Messages = make([]interface{}, 0)
	return
}

func (q *FakeQueueDispatcher) DispatchMessage(message interface{}) (err error) {
	q.Messages = append(q.Messages, message)
	return
}
