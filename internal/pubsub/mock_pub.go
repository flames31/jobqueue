package pubsub

type MockPublisher struct {
	Published []JobEvent
}

func NewMockPublisher() *MockPublisher {
	return &MockPublisher{
		Published: make([]JobEvent, 0),
	}
}

func (mp *MockPublisher) Publish(event JobEvent) error {
	mp.Published = append(mp.Published, event)
	return nil
}
