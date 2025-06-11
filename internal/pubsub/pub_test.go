package pubsub

import (
	"fmt"
	"testing"
	"time"
)

func TestPublisher(t *testing.T) {
	mockPub := NewMockPublisher()

	mockPub.Publish(JobEvent{
		Type:    EventJobCreated,
		JobID:   1,
		UserID:  2,
		Payload: map[string]interface{}{"testing Job number": 1},
		Time:    time.Now(),
	})

	if len(mockPub.Published) != 1 {
		t.Errorf("expected one event to be published")
	}

	fmt.Println(mockPub)
}
